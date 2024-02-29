package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	storages "monopoly-auth/internal/storage"
	"monopoly-auth/tools"
	"time"
)

var (
	ErrPlayerNotFound          = errors.New("player not found")
	ErrEmailOrPasswordNotMatch = errors.New("email or password not match")
)

type Manager interface {
	CreateUser(email string, password string) error

	// FindByEmail(email string) (*models.Player, error)
	// ChangeEmail(player *models.Player, newEmail string, token string) error
	// ChangePassword(player *models.Player, password string, newPassword string) error
}

type Authentication interface {
	SignIn(email string, password string) (string, error)
	SignOut(email string) error
}

type AuthenticationService struct {
	storage storages.PlayerStorage
	logger  *logrus.Logger
}

func (as *AuthenticationService) SignIn(email string, password string) (string, error) {
	player, err := as.storage.FindPlayerByEmail(email)
	if err != nil {
		return "", ErrPlayerNotFound
	}

	hashPassword, err := tools.HashPassword(password)
	if err != nil {
		return "", err
	}

	if !tools.CompareHash(player.PasswordHash, hashPassword) {
		return "", ErrEmailOrPasswordNotMatch
	}

	// TODO: реализация сравнения хэшей паролей и генерация jwt

	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.RegisteredClaims{
			Subject:   player.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		})

	return token.SignedString([]byte(password))
}

func (as *AuthenticationService) SignOut(email string) error {
	return errors.New("not implemented")
}

func NewAuthenticationService(storage storages.PlayerStorage, logger *logrus.Logger) *AuthenticationService {
	return &AuthenticationService{
		storage: storage,
		logger:  logger,
	}
}
