package graphql_pkg

import (
	"context"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	"github.com/graphql-go/graphql"
)

type (
	ResolverRoot interface {
		TestDataResolver() TestDataResolver
	}

	TestDataResolver struct {
		GetAllTestDatas func(ctx context.Context, obj interface{}, next graphql.FieldResolver) (res []*responses.TestRecordList, customErr *apperror.CustomError)
	}

	Resolver struct{}
)
