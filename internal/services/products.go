package services

import (
	"context"
	"fmt"
	"github.com/1makarov/go-csv-crud-example/internal/repository"
	"github.com/1makarov/go-csv-crud-example/internal/types"
	"strings"
)

type ServiceProducts struct {
	repo *repository.RepoProducts
}

func initProducts(repo *repository.RepoProducts) *ServiceProducts {
	return &ServiceProducts{repo: repo}
}

func (s *ServiceProducts) Get(ctx context.Context) (*strings.Reader, error) {
	products, err := s.repo.Get(ctx)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("no products")
	}

	var builder strings.Builder

	for _, product := range products {
		w := fmt.Sprintf("%s;%d\n", product.Name, product.Price)

		builder.WriteString(w)
	}

	return strings.NewReader(builder.String()), nil
}

func (s *ServiceProducts) Update(ctx context.Context, id int64, input types.UpdateInput) error {
	return s.repo.Update(ctx, id, input)
}

func (s *ServiceProducts) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *ServiceProducts) Create(ctx context.Context, input types.CreateInput) (int64, error) {
	return s.repo.Create(ctx, input)
}
