package v1

import (
	"crypto/rand"
	"encoding/hex"
	"gitlab.com/manifold555112/manifold/lib"
	l "gitlab.com/manifold555112/manifold/server/api/lib"
	t "gitlab.com/manifold555112/manifold/server/api/types"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
)

func Login(c *fiber.Ctx) error {
	cnf, err := lib.GetConfig(nil)
	if err != nil {
		return err
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	requestLoginData := t.LoginData{}
	sampleLoginData := t.LoginData{
		Email:    "",
		Password: "",
	}
	err = c.BodyParser(&requestLoginData)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, &sampleLoginData)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	err = validate.Struct(&requestLoginData)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, &sampleLoginData)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	db, err := lib.GetDb(nil)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, &sampleLoginData)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	_, err = db.Use("manifold", "system")
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, &sampleLoginData)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	userAuth, err := l.AuthenticateUser(db, requestLoginData)
	if err != nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.FAILED_TO_AUTHENTICATE_USER), &sampleLoginData)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if userAuth.Authenticated == nil {
		body, status := l.ApiErroResponse(lib.Ptr("Unexpected authentication error"), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if !*userAuth.Authenticated {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.WRONG_PASSWORD), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	res, err := db.Query("select email,verified from $user;", map[string]interface{}{"user": *userAuth.UserId})
	if err != nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.FAILED_TO_FETCH_USER), &sampleLoginData)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	userSlice := make([]*t.User, 0)
	userSlice = append(userSlice, nil)
	_, err = surrealdb.UnmarshalRaw(res, &userSlice)
	if err != nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.FAILED_TO_UNMARSHAL_USER), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if userSlice[0] == nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.USER_NOT_FOUND), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if !userSlice[0].Verified {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.USER_NOT_FOUND), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	authbytes := make([]byte, 64)
	_, err = rand.Read(authbytes)
	if err != nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.FAILED_TO_GENERATE_SESSION_TOKEN), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	authtoken := hex.EncodeToString(authbytes[:])
	sessionSlice := make([]*t.Session, 0)
	sessionSlice = append(sessionSlice, nil)
	res, err = db.Query((&t.Session{}).QueryByUser(*userAuth.UserId))
	_, err = surrealdb.UnmarshalRaw(res, &sessionSlice)
	if err != nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.FAILED_TO_UNMARSHAL_SESSION), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if sessionSlice[0] == nil {
		session := t.Session{
			Id:        nil,
			User:      userAuth.UserId,
			Token:     &authtoken,
			ExpiresAt: lib.Ptr((time.Time{}.Add(time.Hour * 24))),
		}
		res, err := db.Query(session.CreateForUser())
		if err != nil {
			body, status := l.ApiErroResponse(nil, lib.Ptr(t.FAILED_TO_SAVE_SESSION_TO_DATABASE), nil)
			c.SendStatus(status)
			c.SendString(body)
			return nil
		}
		_, err = surrealdb.UnmarshalRaw(res, &sessionSlice)
		if err != nil {
			body, status := l.ApiErroResponse(nil, lib.Ptr(t.FAILED_TO_UNMARSHAL_SESSION), nil)
			c.SendStatus(status)
			c.SendString(body)
			return nil
		}
	}
	if sessionSlice[0] == nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.FAILED_TO_UNMARSHAL_SESSION), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	_, err = db.Query(sessionSlice[0].UpdateToken(authtoken, nil))
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	_, err = db.Query(sessionSlice[0].UpdateExpiresAt(time.Now().Add(time.Hour*24), nil))
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	//TODO cookie secure should be true
	c.Cookie(&fiber.Cookie{
		Name:        "MANIFOLD_SESSION_TOKEN",
		Value:       authtoken,
		Path:        "/",
		Domain:      cnf.Server.FrontendHost,
		MaxAge:      0,
		Expires:     time.Time{}.Add(time.Hour * 24),
		Secure:      true,
		HTTPOnly:    true,
		SameSite:    "None",
		SessionOnly: true,
	})
	return nil
}
