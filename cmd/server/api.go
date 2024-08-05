package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	"github.com/mess110/shortest-path/internal/connections"
)

func main() {
	// rand.Seed(time.Now().UnixNano())

	app := fiber.New()

	context := context.Background()
	driver, err := connections.NewNeo4jConnection(context)
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close(context)

	app.Get("/", func(c fiber.Ctx) error {
		result, err := connections.Neo4jExecuteQuery(context, driver, "match(all{}) return all", map[string]any{})
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.JSON(result)
	})

	app.Get("/shortest_path/:start/:end", func(c fiber.Ctx) error {
		start := c.Params("start")
		end := c.Params("end")

		if start == "" || end == "" {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": "Start and end are required."})
		}

		query := `
			MATCH
				(a:Node {value: $start}),
				(b:Node {value: $end}),
				p = shortestPath((a)-[:RELATED_TO*]-(b))
			RETURN p;
		`

		result, err := connections.Neo4jExecuteQuery(context, driver, query, map[string]any{
			"start": start,
			"end":   end,
		})
		if err != nil {
			return c.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(result)
	})

	app.Post("/query", func(c fiber.Ctx) error {
		payload := struct {
			Query string `json:"query"`
		}{}

		if err := c.Bind().Body(&payload); err != nil {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": err.Error()})
		}

		if payload.Query == "" {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": "Query is required."})
		}

		uuids := []string{}

		// This replaces {"query": "$uuid,@uuid0"}
		// with {"query":"e72b5bc3-98be-46b1-a154-ce1fb2c24c31,e72b5bc3-98be-46b1-a154-ce1fb2c24c31"}
		for strings.Contains(payload.Query, "$uuid") {
			uuid := uuid.NewString()
			uuids = append(uuids, uuid)
			payload.Query = strings.Replace(payload.Query, "$uuid", uuid, 1)
		}
		for i, uuid := range uuids {
			payload.Query = strings.Replace(payload.Query, fmt.Sprintf("@uuid%d", i), uuid, 1)
		}

		results, err := connections.Neo4jExecuteQuery(context, driver, payload.Query, map[string]any{})
		if err != nil {
			return c.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(map[string]any{
			"query":   payload.Query,
			"results": results,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
