package types

type CreateInput struct {
	Name  string `json:"name" db:"name"`
	Price int64  `json:"price" db:"price"`
}

type UpdateInput struct {
	Name  *string `json:"name" db:"name"`
	Price *int64  `json:"price" db:"price"`
}

type Output struct {
	Name  string `json:"name" db:"name"`
	Price int64  `json:"price" db:"price"`
}
