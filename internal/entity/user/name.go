package user

import "fmt"

const (
	minNameLength = 5
	maxNameLength = 15
)

type Username struct {
	value string
}

func NewName(value string) (Username, error) {
	if len(value) < minNameLength || len(value) > maxNameLength {
		return Username{}, fmt.Errorf("username length must be between %d and %d", minNameLength, maxNameLength)
	}

	return Username{
		value: value,
	}, nil
}
