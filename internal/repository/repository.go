package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Products *RepoProducts
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Products: initProducts(db),
	}
}
