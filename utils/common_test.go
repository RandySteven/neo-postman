package utils_test

import (
	"fmt"
	"github.com/RandySteven/neo-postman/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDetailURL(t *testing.T) {
	prefix := "test"
	id := 1
	host := os.Getenv("APP_HOST")
	expectedResp := fmt.Sprintf("%s/%s/%d", host, prefix, id)

	actualResp := utils.DetailURL(prefix, uint64(id))

	assert.Equal(t, expectedResp, actualResp)
}

func TestJsonString(t *testing.T) {
	t.Run("should return json string", func(t *testing.T) {
		request := map[string]interface{}{
			"email":    "test@test.com",
			"password": "test_1234",
		}
		expectedResp := "{\"email\":\"test@test.com\",\"password\":\"test_1234\"}"
		actualResp, _ := utils.JsonString(request)
		assert.Equal(t, expectedResp, actualResp)
	})
}
