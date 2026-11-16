package user

import (
	"github.com/Wh4tisl0ve/cloud-file-storage-go/internal/entity/user"
)

type Repository interface {
	save(*user.User) error
	getByUsername(username string) *user.User
}

type Hasher interface {
	hash(string) (string, error)
	compare(string, string) error
}

type Service struct {
	r Repository
	h Hasher
}

func New(r Repository, h Hasher) Service {
	return Service{
		r: r,
		h: h,
	}
}

func (s Service) create(username, password string) error {
	_, err := user.NewPassword(password)
	if err != nil {
		return err
	}

	passwordHash, err := s.h.hash(password)
	if err != nil {
		return err
	}

	u, err := user.New(username, passwordHash)
	if err != nil {
		return err
	}

	err = s.r.save(u)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) authorize(username, password string) (string, error) {
	_, err := s.h.hash(password)
	if err != nil {
		return "", nil
	}

	u := s.r.getByUsername(username)
	if u == nil {
		return "", nil
	}

	// todo доделать авторизацию
	return "", nil
}
