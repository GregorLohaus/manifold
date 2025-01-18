package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Options(c *fiber.Ctx) error {
	fmt.Println(c.Method())
	if c.Method() == "OPTIONS" {
		fmt.Println("Status Accepted")
		fmt.Printf("%v", &c.Response().Header)
		c.SendStatus(fiber.StatusAccepted)
		return nil
	}
	return c.Next()
}
