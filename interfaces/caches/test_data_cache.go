package caches_interfaces

import "github.com/RandySteven/neo-postman/entities/models"

type TestDataCache interface {
	Cache[models.TestData]
}
