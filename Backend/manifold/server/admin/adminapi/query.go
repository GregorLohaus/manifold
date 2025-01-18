package adminapi

import (
	"encoding/json"
	"fmt"
	"gitlab.com/manifold555112/manifold/lib"

	"github.com/gofiber/fiber/v2"
)

type (
	QueryBody struct {
		QueryString string
	}
)

func Query(c *fiber.Ctx) error {
	db, err := lib.GetDb(nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	body := QueryBody{}
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}
	res, err := db.Query(body.QueryString, nil)
	if err != nil {
		return c.SendString(err.Error())
	}
	resStr, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return err
	}
	return c.SendString(string(resStr))
}
