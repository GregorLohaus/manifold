package v1

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
	"gitlab.com/manifold555112/manifold/lib"
	l "gitlab.com/manifold555112/manifold/server/api/lib"
	t "gitlab.com/manifold555112/manifold/server/api/types"
)

func Verify(c *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	verificationData := t.Verification{}
	sampleRegistrationKey := ""
	sampleVerificationData := t.Verification{
		Email:           "",
		Password:        "",
		RegistrationKey: &sampleRegistrationKey,
	}
	err := c.BodyParser(&verificationData)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, &sampleVerificationData)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	err = validate.Struct(&verificationData)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, &sampleVerificationData)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	db, err := lib.GetDb(nil)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	_, err = db.Use("manifold", "system")
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	authRes, err := l.AuthenticateUser(db, verificationData)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if authRes.Authenticated == nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.USER_NOT_FOUND), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if authRes.UserId == nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.USER_NOT_FOUND), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if !*authRes.Authenticated {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.WRONG_PASSWORD), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	result, err := db.Query("SELECT * FROM $userid", map[string]interface{}{
		"userid": authRes.UserId,
	})
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	userSlice := make([]*t.User, 1)
	_, err = surrealdb.UnmarshalRaw(result, &userSlice)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if userSlice[0] == nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.USER_NOT_FOUND), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if userSlice[0].RegistrationKey == nil {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.REGISTRATION_KEY_NOT_FOUND), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if userSlice[0].Verified {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.USER_ALREADY_VERIFIED), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	if verificationData.RegistrationKey == nil {
		verificationData.RegistrationKey = &sampleRegistrationKey
	}
	if *userSlice[0].RegistrationKey != *verificationData.RegistrationKey {
		body, status := l.ApiErroResponse(nil, lib.Ptr(t.REGISTRATION_KEY_WRONG), nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	mailSum := sha1.Sum([]byte(verificationData.Email))
	mailHex := hex.EncodeToString(mailSum[:])
	_, err = db.Query(
		`BEGIN TRANSACTION;
		DEFINE NAMESPACE `+mailHex+`;
		USE NS `+mailHex+`;
		DEFINE USER `+mailHex+` ON NAMESPACE PASSWORD "`+verificationData.Password+`" ROLES EDITOR;
		DEFINE DATABASE default_workspace;
		USE DB default_workspace;
		DEFINE TABLE schemas SCHEMALESS;
		DEFINE TABLE env SCHEMALESS;
		DEFINE TABLE graphs SCHEMALESS;
		DEFINE TABLE records SCHEMALESS;
		COMMIT;`, nil)
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	_, err = db.Use("manifold", "system")
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	_, err = db.Query("UPDATE $user SET verified = $verified;", map[string]interface{}{
		"user":     *authRes.UserId,
		"verified": true,
	})
	if err != nil {
		body, status := l.ApiErroResponse(lib.Ptr(err.Error()), nil, nil)
		c.SendStatus(status)
		c.SendString(body)
		return nil
	}
	return nil
}
