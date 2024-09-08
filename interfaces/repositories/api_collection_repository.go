package repositories_interfaces

import "github.com/RandySteven/neo-postman/entities/models"

type ApiRepository interface {
	Repositories[models.Api]
	Transaction[models.Api]
}
