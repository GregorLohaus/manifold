package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/manifold555112/manifold/lib"
)

func Headers(c *fiber.Ctx) error {
	cnf, err := lib.GetConfig(nil)
	if err != nil {
		return err
	}
	c.Response().Header.Set("Access-Control-Allow-Origin", cnf.Server.FrontendProtocol+"://"+cnf.Server.FrontendHost+":"+strconv.Itoa(cnf.Server.FrontendPort))
	c.Response().Header.Set("Access-Control-Allow-Credentials", "true")
	c.Response().Header.Set("Access-Control-Allow-Methods", "*")
	c.Response().Header.Set("Access-Control-Allow-Headers", "Content-Type")
	return c.Next()
}
