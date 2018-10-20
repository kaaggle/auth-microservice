package usecase

import (
	"church-adoration/auth"
	"church-adoration/core"

	"church-adoration/models"
)

type authUsecase struct {
	authRepo auth.AuthRepository
}

func NewAuthUsecase(authRepo auth.AuthRepository) auth.AuthRepository {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (a *authUsecase) Login(email, password string) (*models.User, string, error) {
	user, token, err := a.authRepo.Login(email, password)

	if err != nil {
		return nil, "", err
	}

	if user.Email == "" {
		return nil, "", err
	}

	return user, token, nil
}

func (a *authUsecase) Signup(u *models.User) (*models.User, string, error) {
	user, userID, err := a.authRepo.Signup(u)

	if err != nil {
		return nil, "", err
	}

	// send an email with userID to verify
	err = core.SendEmail(
		user.Email,
		"Church Adoration Account Activation",
		"Visit link http://localhost:3300/p/auth/activate/"+userID + " to activate your account.",
		)
	if err != nil {
		return nil, "", err
	}
	return user, userID, nil
}

func (a *authUsecase) ActivateUser(id string) error {
	err := a.authRepo.ActivateUser(id)

	if err != nil {
		return err
	}

	return nil
}
