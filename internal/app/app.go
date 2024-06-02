package app

import (
	"GOncurrently-Calculator/internal/server"
	"context"
	"time"
)

type App struct {
	server *server.Server

	additionDuration       time.Duration
	substractionDuration   time.Duration
	multiplicationDuration time.Duration
	divisionDuration       time.Duration
}

func New(addition, substraction, multiplication, division time.Duration, CP int) *App {
	return &App{
		server: server.New(),

		additionDuration:       addition,
		substractionDuration:   substraction,
		multiplicationDuration: multiplication,
		divisionDuration:       division,
	}
}

func (ap *App) Run() {
	ap.server.Run()
}

func (ap *App) Stop(ctx context.Context) {
	ap.server.Stop(ctx)
}
