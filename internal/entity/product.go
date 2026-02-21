package entity

import (
	"github.com/TobiasTac/go-product-api/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt string    `json:"created_at"`
}
