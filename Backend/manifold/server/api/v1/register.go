package v1

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"gitlab.com/manifold555112/manifold/lib"
	l "gitlab.com/manifold555112/manifold/server/api/lib"
	t "gitlab.com/manifold555112/manifold/server/api/types"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	usr := t.User{}
	sampleUserString := "user:<hex(sha1(email))>"
	sampleExpiery := "2006-01-02T15:04:05Z0700"
	sampleUsr := t.User{
		FirstName:       "",
		LastName:        "",
		Company:         new(string),
		Phone:           new(string),
		Email:           "",
		Password:        "",
		Roles:           []*t.Role{},
		ChildUsers:      []*string{&sampleUserString},
		ParentUser:      &sampleUserString,
		RegistrationKey: nil,
		Verified:        false,
		Plan:            int(t.Trial),
		PlanExpiery:     &sampleExpiery,
	}

	err := c.BodyParser(&usr)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, &sampleUsr)
		c.SendStatus(status)
		c.SendString(body)
		return err
	}
	err = validate.Struct(&usr)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, &sampleUsr)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	db, err := lib.GetDb(nil)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	_, err = db.Use("manifold", "system")
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	registrationKeyByte := make([]byte, 16)
	_, err = rand.Read(registrationKeyByte)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	registrationKey := hex.EncodeToString(registrationKeyByte)
	usr.RegistrationKey = &registrationKey
	emailHash := sha1.Sum([]byte(usr.Email))
	res, _ := db.Select("user:" + hex.EncodeToString(emailHash[:]))
	if res != nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.USER_ALREADY_EXISTS), &sampleUsr)
		c.SendStatus(status)
		c.SendString(body)
		return nil

	}
	fmt.Printf("%v\n", res)
	_, err = db.Create("user:"+hex.EncodeToString(emailHash[:]), &usr)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, &sampleUsr)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	c.SendStatus(200)
	return nil
}
