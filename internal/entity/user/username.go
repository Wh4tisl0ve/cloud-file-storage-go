package user

import (
	"errors"
	"strings"
)

const (
	MIN_LENGTH = 5
	MAX_LENGTH = 15
)

type Username struct {
	value string
}

func NewUserName(v string) (*Username, error) {
	if err := validate(strings.TrimSpace(v)); err != nil {
		return nil, err
	}

	return &Username{
		value: v,
	}, nil
}

func (u Username) String() string {
	return u.value
}

func validate(v string) error {
	if len(v) < MIN_LENGTH {
		return errors.New("username too short")
	}

	if len(v) > MAX_LENGTH {
		return errors.New("username too long")
	}

	return nil
}
