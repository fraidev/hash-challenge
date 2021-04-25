package repositories

import (
	"testing"

	"github.com/fraifelipe/hash-challenge/backend/discounts/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestProductRepository(t *testing.T) {
	db := infrastructure.GetDatabasConnection()
	repository := NewProductRepository(db)

	ids := []uuid.UUID{
		uuid.FromStringOrNil("f0b19f39-b0a8-4b39-a42e-6d83f5cbd4aa"),
		uuid.FromStringOrNil("827bd2cc-9537-4330-ad2b-40885344b71a"),
		uuid.FromStringOrNil("8bf94cd1-98db-44dd-b82a-ce7af028b677"),
		uuid.FromStringOrNil("3d2d1916-e11c-4b87-a514-97b082632b0b"),
	}

	products, err := repository.FindProductByIds(ids)

	if err != nil {
		t.Fatalf("Not expected error")
	}

	if len(products) != 4 {
		t.Fatalf("expected 4 products")
	}

	if products[0].ID.String() != "f0b19f39-b0a8-4b39-a42e-6d83f5cbd4aa" &&
		products[0].PriceInCents != 470000 &&
		products[0].Title != "Playsation 5" &&
		products[0].Description != "Sony video game" {
		t.Fatalf("product not expected")
	}

	if products[1].ID.String() != "827bd2cc-9537-4330-ad2b-40885344b71a" &&
		products[1].PriceInCents != 500000 &&
		products[1].Title != "Xbox One X" &&
		products[1].Description != "Microsoft video game" {
		t.Fatalf("product not expected")
	}

	if products[2].ID.String() != "8bf94cd1-98db-44dd-b82a-ce7af028b677" &&
		products[2].PriceInCents != 200000 &&
		products[2].Title != "Switch" &&
		products[2].Description != "Nintendo video game" {
		t.Fatalf("product not expected")
	}

	if products[3].ID.String() != "3d2d1916-e11c-4b87-a514-97b082632b0b" &&
		products[3].PriceInCents != 500000 &&
		products[3].Title != "PC" &&
		products[3].Description != "A Personal Computer" {
		t.Fatalf("product not expected")
	}
}
