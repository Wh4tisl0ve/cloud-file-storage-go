package app

import (
	"context"
	"fmt"

	"github.com/Wh4tisl0ve/cloud-file-storage-go/cmd/internal/config"
	"github.com/Wh4tisl0ve/cloud-file-storage-go/cmd/pkg/postgres"
)

func Run() {
	c := config.MustLoad()

	dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        c.User,
        c.Password,
        c.Host,
        c.Port,
        c.Name,
    )

	postgres, err := postgres.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(postgres)
}
