package server

import (
	// "encoding/json"
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/manifold555112/manifold/lib"
	"gitlab.com/manifold555112/manifold/server/admin"
	"gitlab.com/manifold555112/manifold/server/api"
	"strconv"
)

func StartServer(args Args) error {
	config, err := lib.GetConfig(args.Conf)
	if err != nil {
		return err
	}
	// confStr, err := json.MarshalIndent(config, "", "  ")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(confStr))
	app := fiber.New()
	admin.Router(args.Conf, app)
	api.Router(args.Conf, app)
	err = app.Listen("0.0.0.0:" + strconv.Itoa(config.Server.Port))
	if err != nil {
		return err
	}
	return nil
}
