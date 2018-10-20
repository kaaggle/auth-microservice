package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Appointment struct {
	ID     bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID bson.ObjectId `json:"user_id" bson:"user_id"`
	User   []User          `json:"user" bson:"user"`

	IsRecursive bool      `json:"is_recursive" bson:"is_recursive"`
	Date        time.Time `json:"date" bson:"date"`
	FromHour        string    `json:"from_hour" bson:"from_hour"`
	ToHour          string    `json:"to_hour" bson:"to_hour"`

	FromMinute        string    `json:"from_minute" bson:"from_minute"`
	ToMinute          string    `json:"to_minute" bson:"to_minute"`

	Cancelled          bool      `json:"cancelled" bson:"cancelled"`
	CancellationReason string    `json:"cancellation_reason" bson:"cancellation_reason"` // if status is cancelled
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
}

type Appointments []Appointment

func (a Appointment) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Date, validation.Required),
		validation.Field(&a.FromHour, validation.Required),
		validation.Field(&a.FromMinute, validation.Required),
		validation.Field(&a.ToHour, validation.Required),
		validation.Field(&a.ToMinute, validation.Required),
	)
}

type CancelAppointment struct {
	CancellationReason string    `json:"cancellation_reason" bson:"cancellation_reason"`
}

func (a CancelAppointment) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.CancellationReason, validation.Required),
	)
}
