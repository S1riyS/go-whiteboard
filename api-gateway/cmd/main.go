package main

import (
	"github.com/S1riyS/go-whiteboard/api-gateway/internal/app"
)

func main() {
	// Setup and run app
	app := app.New()
	app.Run()

}
