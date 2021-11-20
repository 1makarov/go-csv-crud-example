package repository

import (
	"context"
	"fmt"
	"github.com/1makarov/go-csv-crud-example/internal/types"
	"github.com/jmoiron/sqlx"
)

const (
	tableProducts = "products"

	fieldID    = "id"
	fieldName  = "name"
	fieldPrice = "price"

	errNoItemWithThisID = "no item with this id"
	errEmptyRequest     = "empty request"
)

type RepoProducts struct {
	db *sqlx.DB
}

func initProducts(db *sqlx.DB) *RepoProducts {
	return &RepoProducts{db: db}
}

func (r *RepoProducts) Get(ctx context.Context) ([]types.Output, error) {
	var products []types.Output

	if err := r.db.SelectContext(ctx, &products, `select name, price from products`); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *RepoProducts) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `delete from products where id = $1`, id)
	if err != nil {
		return err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if i == 0 {
		return fmt.Errorf(errNoItemWithThisID)
	}

	return nil
}

func (r *RepoProducts) Update(ctx context.Context, id int64, input types.UpdateInput) error {
	var query string

	if input.Name != nil {
		query += fmt.Sprintf("%s = '%s',", fieldName, *input.Name)
	}
	if input.Price != nil {
		query += fmt.Sprintf("%s = '%d' ", fieldPrice, *input.Price)
	}

	if len(query) == 0 {
		return fmt.Errorf(errEmptyRequest)
	}

	req := fmt.Sprintf("update %s set %s where %s = '%d'", tableProducts, query[:len(query)-1], fieldID, id)

	result, err := r.db.ExecContext(ctx, req)
	if err != nil {
		return err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if i == 0 {
		return fmt.Errorf(errNoItemWithThisID)
	}

	return nil
}

func (r *RepoProducts) Create(ctx context.Context, input types.CreateInput) (int64, error) {
	var id int64

	if err := r.db.GetContext(ctx, &id, `

	insert into products (name, price) values ($1, $2) returning id
	
	`, input.Name, input.Price); err != nil {
		return 0, err
	}

	return id, nil
}
