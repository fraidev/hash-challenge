package repositories

import (
	"github.com/fraifelipe/hash-challenge/backend/discounts/models"
	"github.com/go-pg/pg"
	"github.com/satori/go.uuid"
)

type UserRepotory struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepotory {
	return &UserRepotory{db}
}

func (userRepo *UserRepotory) FindUserById(db *pg.DB, id uuid.UUID) (*models.User, error) {
	users := make([]*models.User, 1)

	var err = db.Model(&users).Where("id = (?)", id).Select()

	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return users[0], nil
}