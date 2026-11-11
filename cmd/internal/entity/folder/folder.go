package folder

import (
	"fmt"
	"uuid"
)

type Folder struct {
	id       uuid.UUID
	parentId uuid.UUID
	name     Name
	deleted  bool
}

func NewFolder(parentId uuid.UUID, nameStr string) (*Folder, error) {
	name, err := NewName(nameStr)

	if err != nil {
		return nil, err
	}

	return &Folder{
		id:       uuid.New(),
		parentId: parentId,
		name:     name,
		deleted:  false,
	}, nil
}

func (f *Folder) Rename(newName string) error {
	name, err := NewName(newName)

	if f.name.Equals(name) {
		return fmt.Errorf("folder already has this name")
	}

	if err != nil {
		return err
	}

	f.name = name

	return nil
}

func (f *Folder) Move(parentId uuid.UUID) error {
	if f.parentId == parentId {
		return fmt.Errorf("folder is already in this location")
	}

	f.parentId = parentId

	return nil
}

func (f *Folder) Delete() error {
	if f.deleted {
		return fmt.Errorf("folder is already deleted")
	}

	f.deleted = true

	return nil
}
