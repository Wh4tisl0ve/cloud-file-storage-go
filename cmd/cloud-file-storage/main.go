package main

import (
	"github.com/Wh4tisl0ve/cloud-file-storage-go/internal/app"
	"github.com/Wh4tisl0ve/cloud-file-storage-go/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.Run(&cfg)

}
