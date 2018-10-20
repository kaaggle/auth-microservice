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

func (a *authUsecase) SchoolRegistration(s *models.School) (*models.School, error) {
	school, err := a.authRepo.SchoolRegistration(s)

	if err != nil {
		return nil, err
	}

	return school, nil
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

func (a *authUsecase) Signup(u *models.User) (*models.User, error) {
	user, err := a.authRepo.Signup(u)

	if err != nil {
		return user, err
	}

	return user, nil
}
