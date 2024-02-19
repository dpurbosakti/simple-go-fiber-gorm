package queries

import (
	"gorm.io/gorm"
)

type Querier interface {
	UserQuerier
}

type Queries struct {
	*gorm.DB
}

func NewQueries(db *gorm.DB) Querier {
	return &Queries{DB: db}
}

// var _ Querier = (*Queries)(nil)
