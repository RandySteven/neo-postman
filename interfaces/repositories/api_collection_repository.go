package repositories_interfaces

import "github.com/RandySteven/neo-postman/entities/models"

type ApiCollectionRepository interface {
	Repositories[models.ApiCollection]
}
