package services

import "github.com/1makarov/go-csv-crud-example/internal/repository"

type Services struct {
	Products *ServiceProducts
}

func New(repo *repository.Repository) *Services {
	return &Services{Products: initProducts(repo.Products)}
}
