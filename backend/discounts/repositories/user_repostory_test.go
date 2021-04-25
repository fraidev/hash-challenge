package repositories

import (
	"testing"

	"github.com/fraifelipe/hash-challenge/backend/discounts/infrastructure"
	"github.com/fraifelipe/hash-challenge/backend/discounts/models"
	uuid "github.com/satori/go.uuid"
)

func TestUserRepository(t *testing.T) {
	db := infrastructure.GetDatabasConnection()
	repository := NewUserRepository(db)

	ids := []uuid.UUID{
		uuid.FromStringOrNil("96079cd0-73a7-480a-a8bb-b6bffa0aedb7"),
		uuid.FromStringOrNil("2b6719be-f4b8-4bc4-8989-13f0c4d0d9b8"),
		uuid.FromStringOrNil("9f4abbe9-113c-4099-8bef-2184d030f2c3"),
	}

	users := []models.User{}

	for _, id := range ids {

		user, err := repository.FindUserById(id)
		if err != nil {
			t.Fatalf("Not expected error")
		}

		users = append(users, *user)
	}

	if len(users) != 3 {
		t.Fatalf("expected 3 users")
	}

	if users[0].ID.String() != "96079cd0-73a7-480a-a8bb-b6bffa0aedb7" &&
		users[0].FirstName != "Foo" &&
		users[0].LastName != "Bar" &&
		users[0].DateOfBirth.Format("2006-01-02") != "2001-11-22" {
		t.Fatalf("user not expected")
	}

	if users[1].ID.String() != "2b6719be-f4b8-4bc4-8989-13f0c4d0d9b8" &&
		users[1].FirstName != "Bar" &&
		users[1].LastName != "Foo" &&
		users[1].DateOfBirth.Format("2006-01-02") != "1993-04-10" {
		t.Fatalf("user not expected")
	}

	if users[2].ID.String() != "9f4abbe9-113c-4099-8bef-2184d030f2c3" &&
		users[2].FirstName != "Felipe" &&
		users[2].LastName != "Cardozo" &&
		users[2].DateOfBirth.Format("2006-01-02") != "2021-04-24" {
		t.Fatalf("user not expected")
	}
}
