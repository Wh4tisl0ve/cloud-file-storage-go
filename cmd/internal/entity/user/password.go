package user

import "fmt"

const (
	minPasswordLength = 7
	maxPasswordLength = 50
)

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	if len(value) < minNameLength || len(value) > maxNameLength {
		return Password{}, fmt.Errorf("password length must be between %d and %d", minNameLength, maxNameLength)
	}

	return Password{
		value: value,
	}, nil
}
