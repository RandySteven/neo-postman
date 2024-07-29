package repositories_test

import (
	"context"
	"encoding/json"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/enums"
	"github.com/RandySteven/neo-postman/pkg/postgres"
	"github.com/RandySteven/neo-postman/repositories"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	t.Run("success to save", func(t *testing.T) {
		ctx := context.Background()
		db, _ := postgres.TestDB()
		testDataRepo := repositories.NewTestDataRepository(db)
		testData := &models.TestData{
			Method:               "POST",
			Host:                 os.Getenv("HOST_ENV_VARIABLE"),
			URI:                  "/test",
			Description:          "test mock",
			RequestHeader:        json.RawMessage(`{"Content-Type": "application/json"}`),
			RequestBody:          json.RawMessage(`{"test":"data"}`),
			ExpectedResponseCode: 201,
			ActualResponseCode:   201,
			ExpectedResponse:     json.RawMessage(`{"test":"data"}`),
			ResultStatus:         enums.Expected,
		}

		testData, _ = testDataRepo.Save(ctx, testData)
		assert.NotEqual(t, 0, testData.ID)
	})
}
