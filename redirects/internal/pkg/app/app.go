package app

import (
	"fmt"
	"github.com/alikhanturusbekov/redirects/internal/app/cache"
	"github.com/alikhanturusbekov/redirects/internal/app/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	fiber *fiber.App
}

func New() (*App, error) {
	app := &App{}

	database.ConnectDb()
	cache.ConnectCache()

	app.fiber = fiber.New()

	setupRoutes(app.fiber)

	return app, nil
}

func (a *App) Run() error {
	fmt.Println("server running")

	err := a.fiber.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
