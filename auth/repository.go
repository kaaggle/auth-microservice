package auth

import "schoolsystem/auth-microservice/models"

type AuthRepository interface {
	SchoolRegistration(*models.School) (*models.School, error)
	Login(email, password string) (bool, string, error)
	Signup(*models.User) (*models.User, error)
}
