package usecase

import (
	"schoolsystem/auth-microservice/auth"
	"schoolsystem/auth-microservice/models"
)

type authUsecase struct {
	authRepo auth.AuthRepository
}

func NewAuthUsecase(authRepo auth.AuthRepository) auth.AuthRepository {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (a *authUsecase) Login(email, password string) (bool, error) {
	exists, err := a.authRepo.Login(email, password)

	if err != nil {
		return false, err
	}

	if !exists {
		return false, nil
	}

	return true, nil
}

func (a *authUsecase) Signup(u *models.User) (*models.User, error) {
	user, err := a.authRepo.Signup(u)

	if err != nil {
		return user, err
	}

	return user, nil
}
