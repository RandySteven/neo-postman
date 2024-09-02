package stores

import "github.com/RandySteven/neo-postman/entities/models"

type ApiContentDetailStore interface {
	MongoStore[models.ApiContentDetail]
}
