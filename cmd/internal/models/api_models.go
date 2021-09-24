package models

import (
	"errors"
	"regexp"
)

type UserRegistrationResponse struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}

type EmailRegistrationRequest struct {
	Email string `json:"email"`
}

func (e EmailRegistrationRequest) Validate() error {
	if e.Email == "" {
		return errors.New("Empty email")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(e.Email) {
		return errors.New("Invalid email")
	}

	return nil
}
