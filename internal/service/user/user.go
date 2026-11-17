package user

import (
	"fmt"
	"uuid"

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

func (s Service) createUser(username, password string) error {
	p, err := user.NewPassword(password)
	if err != nil {
		return err
	}

	passwordHash, err := s.h.hash(p.Value())
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

func (s Service) authenticateUser(username, password string) (uuid.UUID, error) {
	u := s.r.getByUsername(username)
	if u == nil {
		return uuid.New(), fmt.Errorf("user with username %s not found", username)
	}

	passwordHash, err := s.h.hash(password)
	if err != nil {
		return uuid.New(), err
	}

	if err := s.h.compare(passwordHash, u.PasswordHash()); err != nil {
		return uuid.New(), fmt.Errorf("incorrect password")
	}

	return u.Id(), nil
}
