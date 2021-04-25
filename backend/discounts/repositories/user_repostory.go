package repositories

import (
	"github.com/fraifelipe/hash-challenge/backend/discounts/models"
	"github.com/go-pg/pg"
	"github.com/satori/go.uuid"
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{db}
}

func (userRepository *UserRepository) FindUserById(id uuid.UUID) (*models.User, error) {
	users := make([]*models.User, 1)

	var err = userRepository.db.Model(&users).Where("id = (?)", id).Select()

	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return users[0], nil
}