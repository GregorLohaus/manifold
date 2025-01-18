package v1

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
	"gitlab.com/manifold555112/manifold/server/api/middleware"
)

func GraphList(c *fiber.Ctx) error {
	if db, ok := c.Locals(middleware.USER_DB_KEY).(*surrealdb.DB); ok {
		res, err := db.Query("select meta::id(id) as id from graphs", nil)
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return nil
		}
		fmt.Printf("%v", res)
		var ids []struct {
			Id string `json:"id"`
		}
		_, err = surrealdb.UnmarshalRaw(res, &ids)
		if err != nil {
			fmt.Println(err.Error())
			c.SendStatus(fiber.StatusInternalServerError)
			return nil
		}
		for _, v := range ids {
			fmt.Println(v)
		}
		jsonres, err := json.Marshal(&ids)
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return nil
		}
		c.SendString(string(jsonres))
		c.SendStatus(fiber.StatusAccepted)
		return nil
	}
	c.SendStatus(fiber.StatusUnauthorized)
	return nil
}
