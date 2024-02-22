package internal

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type PlayerManager struct {
}

func (p *PlayerManager) CreatePlayer(email string, nickname string, password string) (*Player, error) {
	var vErrors []error

	e, err := NewEmail(email)
	if err != nil {
		vErrors = append(vErrors, err)
	}
	n, err := NewNickname(nickname)
	if err != nil {
		vErrors = append(vErrors, err)
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		vErrors = append(vErrors, err)
	}

	if len(vErrors) != 0 {
		return nil, errors.Join(vErrors...)
	}
	return NewPlayer(e, n, passwordHash), nil
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{}
}
