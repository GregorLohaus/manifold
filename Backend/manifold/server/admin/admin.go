package admin

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/manifold555112/manifold/lib"
	"gitlab.com/manifold555112/manifold/server/admin/adminapi"
	"gitlab.com/manifold555112/manifold/server/admin/adminviews"
)

func Router(confPath *string, app *fiber.App) {
	admin := app.Group("/admin")
	admin.Use(adminmocauth)
	adminApi := admin.Group("/api")
	adminapi.Router(confPath, adminApi)
	admin.Get("/", lib.WrapTempl(adminviews.Index))
	admin.Get("/login", lib.WrapTempl(adminviews.Index))
	admin.Get("/query", lib.WrapTempl(adminviews.Query))
	admin.Get("/infotree", lib.WrapTempl(adminviews.Infotree))
	// admin.Get("/login", lib.WrapTempl(adminviews.Login))
	admin.Get("/*", lib.FourOFour())
}
