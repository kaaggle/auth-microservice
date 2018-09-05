package models

import (
	"regexp"
	"time"

	"github.com/globalsign/mgo/bson"
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Name      string        `json:"name"`
	UserID    string        `json:"user_id" bson:"user_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`

	Role string `json:"role" bson:"role"`
}

type Users []User

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email,
			validation.Required.Error("Email is required"),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"))),

		validation.Field(&u.Password, validation.Required, validation.Length(5, 20)),
	)
}
