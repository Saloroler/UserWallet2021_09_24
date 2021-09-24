package registration

import (
	"UserWallet2021_09_24/cmd/internal/db/mysql/repo"
	"UserWallet2021_09_24/cmd/internal/models"
	"UserWallet2021_09_24/cmd/pkg/logger"
	"crypto/rand"
	"fmt"
)

type Process struct {
	userRepo  repo.UserRepo
	tokenRepo repo.TokenRepo
}

func NewProcess(userRepo repo.UserRepo, tokenRepo repo.TokenRepo) Process {
	return Process{userRepo: userRepo, tokenRepo: tokenRepo}
}

func (p *Process) NewUser(email, hash string) (models.User, error) {
	createdUser, err := p.userRepo.Create(models.User{
		Email: email,
		Hash:  hash,
	})

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (p *Process) CreateAuthTokenForUser(user models.User) (models.Token, error) {
	token, err := p.generateToken()
	if err != nil {
		return models.Token{}, err
	}

	newToken := models.Token{
		Token:  token,
		UserID: user.ID,
	}

	err = p.tokenRepo.Create(newToken)
	if err != nil {
		return models.Token{}, err
	}

	return newToken, nil
}

func (p *Process) generateToken() (string, error) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		logger.LogError("Failed to generate simple token, err ", err)
		return "", err
	}
	return fmt.Sprintf("%x", b), nil

}
