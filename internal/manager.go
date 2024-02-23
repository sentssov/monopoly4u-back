package internal

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type PlayerManager struct {
	DB *sqlx.DB
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

	player := NewPlayer(e, n, passwordHash)
	_, err = p.DB.Query("INSERT INTO players (id, email, nickname, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)",
		player.ID, player.Email, player.Nickname, player.PasswordHash, player.CreatedAt)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func NewPlayerManager(db *sqlx.DB) *PlayerManager {
	return &PlayerManager{
		DB: db,
	}
}
