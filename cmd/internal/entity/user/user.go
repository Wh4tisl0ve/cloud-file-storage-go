package user

import (
	"time"
	"uuid"
)

type User struct {
	id        uuid.UUID
	name      Name
	password  Password
	createdAt time.Time
}

func (un *User) New(username, passwordStr string) (*User, error) {
	name, err := NewName(username)
	if err != nil {
		return nil, err
	}

	password, err := NewPassword(passwordStr)
	if err != nil {
		return nil, err
	}

	return &User{
		id:       uuid.New(),
		name:     name,
		password: password,
	}, nil
}
