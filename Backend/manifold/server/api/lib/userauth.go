package lib

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/surrealdb/surrealdb.go"
)

type userAuth interface {
	GetMail() string
	GetPass() string
}

type userAuthResult struct {
	UserId        *string `json:"id,omitempty"`
	Authenticated *bool   `json:"pass_valid,omitempty"`
}

func AuthenticateUser(db *surrealdb.DB, ua userAuth) (*userAuthResult, error) {
	query := "SELECT id,crypto::argon2::compare(password, $password) AS pass_valid FROM $userid"
	email := ua.GetMail()
	mailHash := sha1.Sum([]byte(email))
	userId := "user:" + hex.EncodeToString(mailHash[:])
	fmt.Println(email)
	fmt.Println(userId)
	fmt.Println(ua.GetPass())
	result, err := db.Query(query, map[string]interface{}{
		"userid":   userId,
		"password": ua.GetPass(),
	})
	fmt.Println(result)
	if err != nil {
		return nil, err
	}
	userAuthRes := make([]userAuthResult, 1)
	_, err = surrealdb.UnmarshalRaw(result, &userAuthRes)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &userAuthRes[0], nil
}
