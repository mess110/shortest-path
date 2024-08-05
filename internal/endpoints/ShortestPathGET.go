package endpoints

import (
	"github.com/gofiber/fiber/v3"

	"github.com/mess110/shortest-path/internal/connections/graphdb"
)

func ShortestPathGET(c fiber.Ctx) error {
	db := c.Locals("db").(*graphdb.GraphDB)

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

	result, err := db.ExecuteQuery(query, map[string]any{
		"start": start,
		"end":   end,
	})
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}
