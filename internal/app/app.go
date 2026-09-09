package app

import (
	"fmt"

	"context"

	"github.com/Wh4tisl0ve/cloud-file-storage-go/internal/config"
	"github.com/Wh4tisl0ve/cloud-file-storage-go/pkg/postgres"
)

func Run(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBConfig.Username,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Name,
	)
	postgres := postgres.New(context.Background(), dsn)
	fmt.Println(postgres)

	/* 	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	    w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r) */
	// todo init DB, Cache, S3, logger
}
