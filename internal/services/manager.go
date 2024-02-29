package services

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"monopoly-auth/internal/models"
	storages "monopoly-auth/internal/storage"
)

type PlayerManager struct {
	storage storages.PlayerStorage
	logger  *logrus.Logger
}

func (p *PlayerManager) CreatePlayer(email string, nickname string, password string) error {
	var vErrors []error

	e, err := models.NewEmail(email)
	if err != nil {
		vErrors = append(vErrors, err)
	}
	n, err := models.NewNickname(nickname)
	if err != nil {
		vErrors = append(vErrors, err)
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		vErrors = append(vErrors, err)
	}

	if len(vErrors) != 0 {
		return errors.Join(vErrors...)
	}

	player := models.NewPlayer(e, n, passwordHash)

	if err = p.storage.CreatePlayer(player); err != nil {
		p.logger.Errorf("Error of creating player: %s", err.Error())
	}
	p.logger.Infof("Player (%s) was created with ID %s", player.Email, player.ID)

	return nil
}

func NewPlayerManager(storage storages.PlayerStorage, logger *logrus.Logger) *PlayerManager {
	return &PlayerManager{
		storage: storage,
		logger:  logger,
	}
}
