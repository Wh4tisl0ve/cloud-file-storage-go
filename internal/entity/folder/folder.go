package folder

import (
	"time"
	"uuid"
)

type Folder struct {
	ID         uuid.UUID
	Name       Name
	ParentID   uuid.UUID
	UserID     uuid.UUID
	Deleted    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func New() *Folder {
	// todo invariants and create object
	return nil
}

func (f *Folder) Delete() error {
	if f.Deleted {
		// todo уже удалено
		return nil
	}

	f.Deleted = true
	f.DeletedAt = time.Now()

	return nil
}

func (f *Folder) Rename(newName Name) error {
	if f.Name == newName {
		// todo уже такое имя
		return nil
	}

	f.Name = newName
	f.UpdatedAt = time.Now()

	return nil
}

func (f *Folder) Move(newParentID uuid.UUID) error {
	if f.ParentID == newParentID {
		// todo уже такой родитель
		return nil
	}

	f.ParentID = newParentID
	f.UpdatedAt = time.Now()

	return nil
}
