package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *Login) Validate() error {
	return validation.ValidateStruct(l,
		validation.Field(&l.Email,
			validation.Required.Error("Email is required"),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"))),

		validation.Field(&l.Password, validation.Required, validation.Length(5, 20)),
	)
}
