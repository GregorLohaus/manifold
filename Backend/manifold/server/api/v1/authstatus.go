package v1

import "github.com/gofiber/fiber/v2"

// The [server/api/middleware/UserAuth] middleware gets executed before this
// therefore we can just return AuthStatus 200 before this
func AuthStatus(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusOK)
	return nil
}
