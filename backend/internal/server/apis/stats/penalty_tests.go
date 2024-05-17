package stats

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetPTypes(t *testing.T) {
	// Initialize Fiber app
	app := fiber.New()

	// Register the route
	app.Get("/penaltyTypes", getPTypes)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/penaltyTypes", nil)

	// Create a new HTTP response recorder
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to create test request: %v", err)
	}

	// Assert the status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Assert the response body
	expectedBody := `{"penaltyTypes":["Boarding","Charging","Slashing"]}`
	bodyBytes := make([]byte, resp.ContentLength)
	resp.Body.Read(bodyBytes)
	actualBody := string(bodyBytes)

	assert.JSONEq(t, expectedBody, actualBody)
}

func TestFetchPTypes(t *testing.T) {
	penaltyTypes, err := fetchPTypes()

	// Assert no error
	assert.NoError(t, err)

	// Assert the returned penalty types
	expectedTypes := []string{"Boarding", "Charging", "Slashing"}
	assert.Equal(t, expectedTypes, penaltyTypes)
}
