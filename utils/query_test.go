package utils_test

import (
	"github.com/RandySteven/neo-postman/queries"
	"github.com/RandySteven/neo-postman/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryValidation(t *testing.T) {
	t.Run("the query contain insert", func(t *testing.T) {
		var query queries.GoQuery = `INSERT INTO`
		err := utils.QueryValidation(query, "INSERT")
		assert.Nil(t, err)
	})

	t.Run("the query doesnt insert", func(t *testing.T) {
		var query queries.GoQuery = `SELECT INTO`
		err := utils.QueryValidation(query, "INSERT")
		assert.NotNil(t, err)
	})
}
