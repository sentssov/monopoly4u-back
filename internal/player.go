package internal

import "time"

type Player struct {
	Id           uint      `gorm:"primary_key"`
	Email        string    `gorm:"email"`
	Nickname     string    `bson:"nickname"`
	PasswordHash []byte    `bson:"password"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}
