package storages

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"monopoly-auth/internal/models"
	"time"
)

type PlayerStorage interface {
	FindPlayerByID(id string) (*models.Player, error)
	FindPlayerByEmail(email string) (*models.Player, error)

	CreatePlayer(player *models.Player) error
	UpdatePlayer(player *models.Player) error
	RemovePlayer(player *models.Player) error
}

type SQLPlayerStorage struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

func (s *SQLPlayerStorage) FindPlayerByID(id string) (*models.Player, error) {
	rows, err := s.db.Query("SELECT * FROM players WHERE id = $1", id)
	if err != nil {
		s.logger.Errorf("Player with specified id was not found: %s", id)
		return nil, err
	}

	playerDB := &PlayerDB{}
	err = rows.Scan()
	if err != nil {
		return nil, err
	}

	pId, _ := models.NewID(playerDB.ID)
	email, _ := models.NewEmail(playerDB.Email)
	nickname, _ := models.NewNickname(playerDB.Nickname)

	player := &models.Player{
		ID:           pId,
		Email:        email,
		Nickname:     nickname,
		PasswordHash: playerDB.PasswordHash,
		CreatedAt:    playerDB.CreatedAt,
		UpdatedAt:    playerDB.UpdatedAt,
	}

	return player, nil
}

func (s *SQLPlayerStorage) FindPlayerByEmail(email string) (*models.Player, error) {
	return nil, errors.New("not implemented")
}

func (s *SQLPlayerStorage) CreatePlayer(player *models.Player) error {
	_, err := s.db.Query("INSERT INTO players (id, email, nickname, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)",
		player.ID, player.Email, player.Nickname, player.PasswordHash, player.CreatedAt)
	if err != nil {
		s.logger.Errorf("Error inserting player with ID %s", player.ID)
		return err
	}
	return nil
}

func (s *SQLPlayerStorage) UpdatePlayer(player *models.Player) error {
	_, err := s.db.Exec("UPDATE players SET email = $1, nickname = $2, password_hash = $3 WHERE id = $4)",
		player.Email, player.Nickname, player.PasswordHash, player.ID)
	if err != nil {
		s.logger.Errorf("Error updating player with ID %s", player.ID)
		return err
	}
	return nil
}

func (s *SQLPlayerStorage) RemovePlayer(player *models.Player) error {
	_, err := s.db.Exec("DELETE FROM players WHERE id = $1", player.ID)
	if err != nil {
		s.logger.Errorf("Error removing player with ID %s", player.ID)
		return err
	}
	return nil
}

func NewSQLPlayerStorage(db *sqlx.DB, logger *logrus.Logger) *SQLPlayerStorage {
	return &SQLPlayerStorage{
		db:     db,
		logger: logger,
	}
}

type PlayerDB struct {
	ID           string
	Email        string
	Nickname     string
	PasswordHash []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
