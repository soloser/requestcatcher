package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/blackangelnk/requestcatcher/internal/app"
	"github.com/blackangelnk/requestcatcher/internal/config"
	"github.com/blackangelnk/requestcatcher/internal/storage"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.Init()

	s := storage.NewDb(sqlx.MustOpen("sqlite3", ":memory:"))
	app := app.NewApp(cfg, s)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ctx := context.Background()
	app.Run()
	log.Print("App started")
	<-done

	app.Stop(ctx)
	log.Print("App stopped")
}
