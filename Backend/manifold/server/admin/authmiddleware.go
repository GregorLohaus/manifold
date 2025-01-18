package admin

import (
	"fmt"
	"gitlab.com/manifold555112/manifold/server/admin/adminapi"
	"gitlab.com/manifold555112/manifold/server/admin/adminviews"

	"github.com/gofiber/fiber/v2"
)

func adminauth(c *fiber.Ctx) error {
	if adminapi.ADMINSESSION != nil && *adminapi.ADMINSESSION == c.Cookies("ADMINSESSION") {
		c.Locals(adminviews.ADMIN_LOGGED_IN, 1)
	} else {
		c.Locals(adminviews.ADMIN_LOGGED_IN, 0)
		fmt.Println("Route:" + c.Path())
		if c.Path() != "/admin/api/login" && c.Path() != "/admin/login" {
			return c.Redirect("/admin/login")
		}
	}
	auth := c.Get(fiber.HeaderAuthorization)
	fmt.Println(auth)
	return c.Next()
}

func adminmocauth(c *fiber.Ctx) error {
	c.Locals(adminviews.ADMIN_LOGGED_IN, 1)
	return c.Next()
}
