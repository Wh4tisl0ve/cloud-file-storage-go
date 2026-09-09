package user

import (
	"time"
	"uuid"
)

type User struct {
	ID           uuid.UUID
	Username     Username
	PasswordHash string
	CreatedAt    time.Time
}

func New(username Username, pwdHash string) *User {
	return &User{
		ID:           uuid.New(),
		Username:     username,
		PasswordHash: pwdHash,
		CreatedAt:    time.Now(),
	}
}
