package repositories

import (
	"github.com/fraifelipe/hash-challenge/backend/discounts/models"
	"github.com/go-pg/pg"
	"github.com/satori/go.uuid"
)

type ProductRepository struct {
	db *pg.DB
}

func NewProductRepository(db *pg.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (productRepository *ProductRepository) FindProductByIds(ids []uuid.UUID) ([]*models.Product, error) {
	products := make([]*models.Product, len(ids))

	var err = productRepository.db.Model(&products).WhereIn("id IN (?)", ids).Select()

	if err != nil {
		return nil, err
	}

	return products, nil
}
