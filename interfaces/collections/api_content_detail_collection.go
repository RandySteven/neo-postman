package collections_interfaces

import "github.com/RandySteven/neo-postman/entities/models"

type ApiContentDetailCollection interface {
	MongoCollection[models.ApiContentDetail]
}
