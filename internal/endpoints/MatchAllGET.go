package endpoints

import (
	"github.com/gofiber/fiber/v3"

	"github.com/mess110/shortest-path/internal/connections/graphdb"
)

func MatchAllGET(c fiber.Ctx) error {
	db := c.Locals("db").(*graphdb.GraphDB)

	result, err := db.ExecuteQuery("match(all{}) return all", map[string]any{})
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(result)
}
