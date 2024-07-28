package repositories_test

import (
	"context"
	"encoding/json"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/enums"
	"github.com/RandySteven/neo-postman/mocks"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	t.Run("success to save", func(t *testing.T) {
		testDataRepo := new(mocks.TestDataRepository)
		ctx := context.Background()
		testData := &models.TestData{
			Method:               "POST",
			Host:                 os.Getenv("HOST_ENV_VARIABLE"), // Replace with actual environment variable name
			URI:                  "/test",
			Description:          "test mock",
			RequestHeader:        json.RawMessage(`{"Content-Type": "application/json"}`),
			RequestBody:          json.RawMessage(`{"test":"data"}`),
			ExpectedResponseCode: 201,
			ActualResponseCode:   201,
			ExpectedResponse:     json.RawMessage(`{"test":"data"}`),
			ResultStatus:         enums.Expected,
		}

		testDataRepo.On("Save", ctx, testData).Return(testData, nil).Once()

		result, err := testDataRepo.Save(ctx, testData)

		assert.NoError(t, err)
		assert.Equal(t, testData, result)

		testDataRepo.AssertExpectations(t)
	})
}
