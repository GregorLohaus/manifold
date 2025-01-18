package adminapi

import "github.com/gofiber/fiber/v2"

var ADMINSESSION *string

func Router(confPath *string, adminApi fiber.Router) {
	adminApi.Post("/login", Login)
	adminApi.Post("/query", Query)
}
