package file

import (
	"fmt"
	"uuid"
)

type File struct {
	id       uuid.UUID
	parentId uuid.UUID
	name     Name
	deleted  bool
}

func NewFile(parentId uuid.UUID, nameStr string) (*File, error) {
	name, err := NewName(nameStr)

	if err != nil {
		return nil, err
	}

	return &File{
		id:       uuid.New(),
		parentId: parentId,
		name:     name,
		deleted:  false,
	}, nil
}

func (f *File) Rename(newName string) error {
	name, err := NewName(newName)

	if f.name.Equals(name) {
		return fmt.Errorf("file already has this name")
	}

	if err != nil {
		return err
	}

	f.name = name

	return nil
}

func (f *File) Move(parentId uuid.UUID) error {
	if f.parentId == parentId {
		return fmt.Errorf("file is already in this location")
	}

	f.parentId = parentId

	return nil
}

func (f *File) Delete() error {
	if f.deleted {
		return fmt.Errorf("file is already deleted")
	}

	f.deleted = true

	return nil
}
