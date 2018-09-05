package models

import (
	"regexp"
	"time"

	"github.com/globalsign/mgo/bson"
	validation "github.com/go-ozzo/ozzo-validation"
)

type School struct {
	ID              bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name            string        `json:"name"`
	Email           string        `json:"email"`
	PhoneNumber     string        `json:"phone_number" bson:"phone_number"`
	InstitutionName string        `json:"institution_name" bson:"institution_name"`
	Location        string        `json:"location"`
	Website         string        `json:"website" `
	Message         string        `json:"message"`
	Approved        bool          `json:"approved"`
	Type            string        `json:"type"`
	CreatedAt       time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" bson:"updated_at"`
}

func (s *School) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Email,
			validation.Required.Error("Email is required"),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"))),

		validation.Field(&s.InstitutionName, validation.Required),
		validation.Field(&s.Name, validation.Required),
	)
}
