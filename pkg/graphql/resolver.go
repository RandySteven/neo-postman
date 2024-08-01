package graphql_pkg

type Resolver struct {
	TestDataResolver interface {
		CreateTestData()
	}
}
