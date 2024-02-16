package main

import (
	"os"
	"time"

	"github.com/dpurbosakti/fiber-gorm/db"
	"github.com/dpurbosakti/fiber-gorm/route"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	db.DatabaseInit()
	db.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal().Msg("failed to serve")
	}
}
