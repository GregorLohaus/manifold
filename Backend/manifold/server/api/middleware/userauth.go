package middleware

import (
	"fmt"
	"runtime/debug"
	"time"

	"gitlab.com/manifold555112/manifold/lib"
	t "gitlab.com/manifold555112/manifold/server/api/types"

	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
)

const USER_DB_KEY = "USER_DB"

func UserAuth(c *fiber.Ctx) error {
	if c.Path() == "/api/v1/login" || c.Path() == "/api/v1/register" || c.Path() == "/api/v1/verify" {
		return c.Next()
	}
	if c.Cookies("MANIFOLD_SESSION_TOKEN") == "" {
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	db, err := lib.GetDb(nil)
	if err != nil {
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	_, err = db.Use("manifold", "system")
	if err != nil {
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	session := t.Session{}
	sessionSlice := make([]*t.Session, 0)
	sessionSlice = append(sessionSlice, nil)
	query, args := session.QueryByToken(c.Cookies("MANIFOLD_SESSION_TOKEN"))
	res, err := db.Query(query, args)
	if err != nil {
		fmt.Println(err.Error())
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	_, err = surrealdb.UnmarshalRaw(res, &sessionSlice)
	if err != nil {
		fmt.Println(err.Error())
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	if sessionSlice[0] == nil {
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	sessionSlice[0].DBExpieryToExpiery()
	if sessionSlice[0].ExpiresAt.Before(time.Now()) {
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	userSlice := make([]*t.User, 0)
	userSlice = append(userSlice, nil)
	res, err = db.Query(lib.Ptr(t.User{}).Query(sessionSlice[0].User))
	_, err = surrealdb.UnmarshalRaw(res, &userSlice)
	if userSlice[0] == nil {
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	mailHash := userSlice[0].MailHash()
	if userDb, ok := c.Locals(USER_DB_KEY).(*surrealdb.DB); ok {
		if userDb == nil {
			userDb, err = lib.NewDb()
			if err != nil {
				fmt.Println(err.Error())
				debug.PrintStack()
				c.SendStatus(fiber.StatusUnauthorized)
				return nil
			}
			c.Locals(USER_DB_KEY, userDb)
		}
		//TODO store workspace in session and check here
		_, err = userDb.Use(*mailHash, "default_workspace")
		if err != nil {
			fmt.Println(err.Error())
			debug.PrintStack()
			c.SendStatus(fiber.StatusUnauthorized)
			return nil
		}
		return c.Next()
	}
	userDb, err := lib.NewDb()
	if err != nil {
		fmt.Println(err.Error())
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	//TODO store workspace in session and check here
	_, err = userDb.Use(*mailHash, "default_workspace")
	if err != nil {
		fmt.Println(err.Error())
		debug.PrintStack()
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}
	c.Locals(USER_DB_KEY, userDb)
	return c.Next()
}
