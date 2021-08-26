package datastore

import (
	"fmt"

	domain "cart-checkout-simulation/domain"
)

var (
	products = []domain.Product{
		{
			ID:          1,
			Title:       "Ergonomic Wooden Pants",
			Description: "Deleniti beatae porro.",
			Amount:      15157,
			IsGift:      false,
		},
		{
			ID:          2,
			Title:       "Ergonomic Cotton Keyboard",
			Description: "Iste est ratione excepturi repellendus adipisci qui.",
			Amount:      93811,
			IsGift:      false,
		},
		{
			ID:          3,
			Title:       "Gorgeous Cotton Chips",
			Description: "Nulla rerum tempore rem.",
			Amount:      60356,
			IsGift:      false,
		},
		{
			ID:          4,
			Title:       "Fantastic Frozen Chair",
			Description: "Et neque debitis omnis quam enim cupiditate.",
			Amount:      56230,
			IsGift:      false,
		},
		{
			ID:          5,
			Title:       "Incredible Concrete Soap",
			Description: "Dolorum nobis temporibus aut dolorem quod qui corrupti.",
			Amount:      42647,
			IsGift:      false,
		},
		{
			ID:          6,
			Title:       "Handcrafted Steel Towels",
			Description: "Nam ea sed animi neque qui non quis iste.",
			Amount:      900,
			IsGift:      true,
		},
	}
)

type database struct {
	products []domain.Product
}

type Database interface {
	FindAll() ([]domain.Product, error)
	Get(ID int32) (domain.Product, error)
}

func NewDatabase() Database {
	return &database{products}
}

func (db *database) FindAll() ([]domain.Product, error) {
	return db.products, nil
}

func (db *database) Get(ID int32) (domain.Product, error) {
	for _, p := range db.products {
		if p.ID == ID {
			return p, nil
		}
	}

	return domain.Product{}, fmt.Errorf("product not found: %d", ID)
}
