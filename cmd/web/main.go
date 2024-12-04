package main

import (
	"fmt"

	"github.com/adepuu/tokopaedi-backend/internal/config"
)

func main() {
	viperConfig := config.LoadConfig()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	app := config.NewFiber(viperConfig)

	cfg := &config.AppConfig{
		DB:     db,
		App:    app,
		Log:    log,
		Config: viperConfig,
	}

	cfg.Run()

	webPort := viperConfig.GetInt("APP_PORT")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
