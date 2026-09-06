package app

import (
	"fmt"

	"github.com/Wh4tisl0ve/cloud-file-storage-go/internal/config"
)

func Run(cfg *config.Config) {
	fmt.Println(cfg)

	// todo init DB, Cache, S3, logger
}
