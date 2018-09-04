package auth

import "schoolsystem/auth-microservice/models"

type AuthUsecase interface {
	Login(email, password string) (bool, error)
	Signup(*models.User) (*models.User, error)
}
