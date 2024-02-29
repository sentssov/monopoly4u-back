package models

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
	"time"
)

var (
	ErrInvalidID       = errors.New("invalid id provided")
	ErrInvalidEmail    = errors.New("invalid email provided")
	ErrInvalidNickname = errors.New("invalid nickname provided")

	EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+\\.[a-zA-Z]{2,}$")
)

type Player struct {
	ID           ID        `gorm:"primary_key"`
	Email        Email     `gorm:"email"`
	Nickname     Nickname  `bson:"nickname"`
	PasswordHash []byte    `bson:"password"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func NewPlayer(email Email, nickname Nickname, passwordHash []byte) *Player {
	return &Player{
		ID:           ID(uuid.NewString()),
		Email:        email,
		Nickname:     nickname,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
}

type ID string

func NewID(i string) (ID, error) {
	_, err := uuid.Parse(i)
	if err != nil {
		return "", errors.Join(ErrInvalidID)
	}
	return ID(i), nil
}

func (i ID) String() string {
	return string(i)
}

type Email string

func NewEmail(e string) (Email, error) {
	if !EmailRegex.MatchString(e) {
		return "", errors.Join(ErrInvalidEmail)
	}
	return Email(e), nil
}

func (e Email) String() string {
	return string(e)
}

type Nickname string

func NewNickname(n string) (Nickname, error) {
	// Validation rule for nickname
	return Nickname(n), nil
}

func (n Nickname) String() string {
	return string(n)
}
