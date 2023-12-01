package usecase

import (
	"errors"
	"fww-wrapper/internal/repository"
	"fww-wrapper/internal/tools"
)

type usecase struct {
	repository repository.Repository
}

// GenerateToken implements Usecase.
func (u *usecase) GenerateToken(username string, password string) (string, error) {
	user, err := u.repository.FindUserByUsername(username)
	if err != nil {
		return "", err
	}

	hashPassword := tools.HashPassword(password)

	if user.Password != hashPassword {
		return "", errors.New("invalid username and password")
	}

	token, err := tools.GenerateTokenJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

type Usecase interface {
	GenerateToken(username string, password string) (string, error)
}

func NewUsecase(repository repository.Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}
