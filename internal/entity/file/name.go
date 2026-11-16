package file

import "fmt"

const (
	minNameLength = 5
	maxNameLength = 30
)

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	if len(value) < minNameLength || len(value) > maxNameLength {
		return Name{}, fmt.Errorf("file name length must be between %d and %d", minNameLength, maxNameLength)
	}

	return Name{
		value: value,
	}, nil
}

func (n Name) Equals(otherName Name) bool {
	return n.value == otherName.value
}
