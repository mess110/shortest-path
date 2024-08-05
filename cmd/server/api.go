package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/mess110/shortest-path/internal/connections/graphdb"
	"github.com/mess110/shortest-path/internal/endpoints"
)

func main() {
	app := fiber.New()

	context := context.Background()
	db, err := graphdb.NewConnection(context)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app.Use(func(c fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Get("/", endpoints.MatchAllGET)
	app.Get("/shortest-path/:start/:end", endpoints.ShortestPathGET)
	app.Post("/query", endpoints.QueryPOST)

	log.Fatal(app.Listen(":3000"))
}
