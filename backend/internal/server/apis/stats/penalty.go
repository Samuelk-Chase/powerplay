package stats

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/penaltyTypes", auth.Public, getPTypes)
}

func getPTypes(c *fiber.Ctx) error {
	penaltyTypes, err := fetchPTypes()
	if err != nil {
		log.Printf("Error getting penalty types from server: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"penaltyTypes": penaltyTypes})
}

func fetchPTypes() ([]string, error) {
	return []string{"Boarding", "Charging", "Slashing"}, nil
}
