package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/manifold555112/manifold/lib"
	"gitlab.com/manifold555112/manifold/server/api/middleware"
	apiv1 "gitlab.com/manifold555112/manifold/server/api/v1"
)

func Router(confPath *string, app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(middleware.Headers, middleware.Options, middleware.UserAuth)
	v1.Post("/register", apiv1.Register)
	v1.Post("/verify", apiv1.Verify)
	v1.Post("/login", apiv1.Login)
	v1.Get("/authstatus", apiv1.AuthStatus)
	// v1.Get("/graph")
	// v1.Get("/graph/:id")
	v1.Post("/graph/:id", apiv1.PostGraph)
	v1.Get("/graph/:id", apiv1.GraphDetail)
	v1.Get("/graphs", apiv1.GraphList)
	// v1.Get("/schema")
	// v1.Get("/schema/:id")
	// v1.Post("/schema/:id")
	api.Get("/*", lib.FourOFour())
}
