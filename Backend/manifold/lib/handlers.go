package lib

import (
	"errors"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func FourOFour() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return errors.New("Not Found")
	}
}

func WrapTempl(templComponent func(fiberContext *fiber.Ctx) templ.Component) func(fiberContext *fiber.Ctx) error {
	return func(fiberContext *fiber.Ctx) error {
		templHandler := templ.Handler(templComponent(fiberContext))
		handler := fasthttpadaptor.NewFastHTTPHandler(templHandler)
		handler(fiberContext.Context())
		return nil
	}
}
