package auth

import "church-adoration/models"

type AuthRepository interface {
	ActivateUser(id string) error
	Login(email, password string) (*models.User, string, error)
	Signup(*models.User) (*models.User, string, error)
}
