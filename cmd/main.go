package main

import (
	"fmt"
	"log"

	"github.com/galaxy-empire-team/notifier/internal/app"
	"github.com/galaxy-empire-team/notifier/internal/config"
	"github.com/galaxy-empire-team/notifier/internal/db"
	"github.com/galaxy-empire-team/notifier/internal/httpserver"
	notificationservice "github.com/galaxy-empire-team/notifier/internal/service/notification"
	notificationstorage "github.com/galaxy-empire-team/notifier/internal/storage/notification"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("config.New(): %w", err)
	}

	ctx, app, err := app.New(cfg.App)
	if err != nil {
		return fmt.Errorf("app.New(): %w", err)
	}

	// initialize pgx infra
	db, err := db.New(ctx, cfg.PgConn)
	if err != nil {
		return fmt.Errorf("db.New(): %w", err)
	}
	defer db.Close()

	// initialize storages
	notificationStorage := notificationstorage.New(db)

	// initialize services
	notificationService := notificationservice.New(notificationStorage)

	// initialize http server
	httpServer := httpserver.New(app.ComponentLogger("httpserver"))

	httpServer.RegisterRoutes(notificationService)

	err = httpServer.Start(ctx, cfg.Server)
	if err != nil {
		return fmt.Errorf("httpServer.Start(): %w", err)
	}

	return nil
}
