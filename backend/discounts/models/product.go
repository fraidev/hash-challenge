package models

import (
	"github.com/satori/go.uuid"
)

type Product struct {
	tableName    struct{}  `sql:"product"`
	ID           uuid.UUID `json:"id", sql:"type:uuid"`
	PriceInCents uint64    `json:"price_in_cents"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
}