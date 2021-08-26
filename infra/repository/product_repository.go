package repository

import (
	"cart-checkout-simulation/domain"

	datastore "cart-checkout-simulation/infra/datastore"
)

type productRepository struct {
	db datastore.Database
}

type ProductRepository interface {
	FindAll() ([]domain.Product, error)
	Get(ID int32) (domain.Product, error)
}

func NewProductRespository(db datastore.Database) ProductRepository {
	return &productRepository{db}
}

func (pr productRepository) FindAll() ([]domain.Product, error) {
	return pr.db.FindAll()
}

func (pr productRepository) Get(ID int32) (domain.Product, error) {
	return pr.db.Get(ID)
}
