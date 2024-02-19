package handlers

import (
	"github.com/dpurbosakti/fiber-gorm/app/queries"
	"github.com/dpurbosakti/fiber-gorm/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	query     queries.Querier
	validator *validator.Validate
}

func New(query queries.Querier) *Handler {
	validate := utils.NewValidator()
	return &Handler{
		query:     query,
		validator: validate,
	}
}
