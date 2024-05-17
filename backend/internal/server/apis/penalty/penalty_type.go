package penalty

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

type PenaltyType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/penaltyTypes", auth.Public, getPenaltyTypes)
}

// Mock function to simulate fetching penalty types
func fetchPenaltyTypes() ([]PenaltyType, error) {
	// Mock data - replace this with actual database logic if needed
	penaltyTypes := []PenaltyType{
		{ID: 1, Name: "Minor Penalty"},
		{ID: 2, Name: "Major Penalty"},
		{ID: 3, Name: "Misconduct Penalty"},
	}

	// Simulate a potential database error
	if len(penaltyTypes) == 0 {
		return nil, errors.New("no penalty types found")
	}

	return penaltyTypes, nil
}

func getPenaltyTypes(c *fiber.Ctx) error {
	log := locals.Logger(c)

	// Fetch penalty types
	penaltyTypes, err := fetchPenaltyTypes()
	if err != nil {
		log.WithErr(err).Error("Failed to fetch penalty types")
		return responder.InternalServerError(c, "Failed to fetch penalty types")
	}

	// Log the action
	log.Info("Retrieved penalty types")

	// Respond with the penalty types as JSON
	return responder.OkWithData(c, penaltyTypes)
}
