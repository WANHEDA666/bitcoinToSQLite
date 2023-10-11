package main

import (
	"fmt"
	"log/slog"
	"os"
	"testProject/api"
	"testProject/internal/config"
	"testProject/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	storage, err := sqlite.New(cfg.StoragePath)

	if err != nil {
		os.Exit(1)
	}

	var response = api.GetBitcoinData()
	_, err = sqlite.SaveData(response.Time.Updated, response.BPI.USD.Rate, response.BPI.EUR.Rate, response.BPI.GBP.Rate, storage)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
