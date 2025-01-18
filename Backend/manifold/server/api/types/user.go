package types

import (
	"crypto/sha1"
	"encoding/hex"
	"gitlab.com/manifold555112/manifold/lib"
)

type Role int

const (
	Owner Role = iota + 1
	Admin
	Editor
	Viewer
)

type Plan int

const (
	Trial Plan = iota + 1
)

type User struct {
	Id              *string   `json:"id,omitempty"`
	FirstName       string    `json:"first_name" validat:"required"`
	LastName        string    `json:"last_name" validate:"required"`
	Company         *string   `json:"company,omitempty"`
	Phone           *string   `json:"phone,omitempty" validate:"omitnil,e164"`
	Email           string    `json:"email" validate:"required,email"`
	Password        string    `json:"password" validate:"required"`
	Roles           []*Role   `json:"roles" validate:"required"`
	ChildUsers      []*string `json:"child_users,omitempty"`
	ParentUser      *string   `json:"parent_user,omitempty"`
	RegistrationKey *string   `json:"registration_key,omitempty"`
	Verified        bool      `json:"verified"`
	Plan            int       `json:"plan" validate:"required"`
	PlanExpiery     *string   `json:"plan_expiery,omitempty" validate:"omitnil,datetime=2006-01-02T15:04:05Z0700"`
}

func (u *User) Query(userId *string) (query string, args map[string]interface{}) {
	query = `select id,first_name,last_name,company,phone,email,parent_user,registration_key,verified,plan,plan_expiery from $user_id`
	if userId != nil {
		args = map[string]interface{}{
			"user_id": *userId,
		}
		return
	}
	args = map[string]interface{}{
		"user": u.Id,
	}
	return
}

func (u *User) MailHash() (mailhash *string) {
	mailHash := sha1.Sum([]byte(u.Email))
	mailhash = lib.Ptr(hex.EncodeToString(mailHash[:]))
	return
}
