package endpoints

import (
	"github.com/gofiber/fiber/v3"

	"github.com/mess110/shortest-path/internal/connections/graphdb"
	"github.com/mess110/shortest-path/internal/utils"
)

func QueryPOST(c fiber.Ctx) error {
	db := c.Locals("db").(*graphdb.GraphDB)

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

	payload.Query = utils.InjectUUID(payload.Query)

	results, err := db.ExecuteQuery(payload.Query, map[string]any{})
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(map[string]any{
		"query":   payload.Query,
		"results": results,
	})
}
