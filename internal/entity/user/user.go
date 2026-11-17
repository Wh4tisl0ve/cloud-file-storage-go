package user

import (
	"time"
	"uuid"
)

type User struct {
	id           uuid.UUID
	username     Username
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
		username:     name,
		passwordHash: password,
	}, nil
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) Id() uuid.UUID {
	return u.id
}
