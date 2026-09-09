package repository

import (
	"context"

	"github.com/Wh4tisl0ve/cloud-file-storage-go/internal/entity/user"
	"github.com/Wh4tisl0ve/cloud-file-storage-go/pkg/postgres"
)

type UserRepository struct {
	dbConn postgres.Postgres
}

func New(dbConn *postgres.Postgres) UserRepository {
	return UserRepository{
		dbConn: *dbConn,
	}
}

func (ur *UserRepository) Save(u user.User) error {
	query := `INSERT INTO users (
					id,
					username,
					password_hash,
					created_at
				)
				VALUES ($1, $2, $3, $4)`

	_, err := ur.dbConn.Pool.Exec(
		context.Background(),
		query,
		u.ID, u.Username, u.PasswordHash, u.CreatedAt,
	)

	if err != nil {
		// todo более понятные типы ошибок: conflict и тд
		return err
	}

	return nil
}
