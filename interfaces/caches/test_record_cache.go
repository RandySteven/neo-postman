package caches_interfaces

import "github.com/RandySteven/neo-postman/entities/models"

type TestRecordCache interface {
	Cache[models.TestRecord]
}
