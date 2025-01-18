package adminapi

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/manifold555112/manifold/lib"
	"time"
)

func Login(c *fiber.Ctx) error {
	config, err := lib.GetConfig(nil)
	if err != nil {
		return err
	}
	fmt.Println("Adminapi login")
	password := c.FormValue("password")
	passhash := md5.Sum([]byte(password))
	fmt.Println(hex.EncodeToString(passhash[:]))
	fmt.Println(config.Server.AdminPasswordHash)
	if hex.EncodeToString(passhash[:]) == config.Server.AdminPasswordHash {
		fmt.Println("success")
		authbytes := make([]byte, 64)
		_, err := rand.Read(authbytes)
		if err != nil {
			return err
		}
		authtoken := hex.EncodeToString(authbytes[:])
		ADMINSESSION = &authtoken
		c.Cookie(&fiber.Cookie{
			Name:        "ADMINSESSION",
			Value:       authtoken,
			Path:        "",
			Domain:      "",
			MaxAge:      0,
			Expires:     time.Time{},
			Secure:      true,
			HTTPOnly:    true,
			SameSite:    "Lax",
			SessionOnly: true,
		})
	}
	c.Response().Header.Add("HX-Redirect", "/admin")
	return nil
}
