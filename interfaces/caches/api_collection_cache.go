package caches_interfaces

import "github.com/RandySteven/neo-postman/entities/models"

type ApiCollectionCache interface {
	Cache[models.Api]
}
