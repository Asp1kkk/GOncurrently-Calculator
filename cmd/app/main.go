package main

import (
	"GOncurrently-Calculator/internal/app"
)

func main() {
	app := app.New(100, 100, 100, 100, 5)
	app.Run()
}
