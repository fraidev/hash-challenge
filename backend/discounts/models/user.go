package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type User struct {
	tableName    struct{}  `sql:"user"`
	ID           uuid.UUID `json:"id", sql:"type:uuid"`
	FirstName    string    `json:"first_name`
	LastName     string    `json:"last_name"`
	DateOfBirth  time.Time `json:"date_of_birth"`
}