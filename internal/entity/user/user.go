package user

import (
	"time"
	"uuid"
)

type User struct {
	id           uuid.UUID
	name         Name
	passwordHash string
	createdAt    time.Time
}

func New(username, password string) (*User, error) {
	name, err := NewName(username)
	if err != nil {
		return nil, err
	}

	return &User{
		id:           uuid.New(),
		name:         name,
		passwordHash: password,
	}, nil
}
