package app

import (
	repository "cart-checkout-simulation/infra/repository"
	cart_request "cart-checkout-simulation/input/cart/request"
	cart_response "cart-checkout-simulation/input/cart/response"
	product_response "cart-checkout-simulation/input/product/response"
)

type cartApplication struct {
	productRepository repository.ProductRepository
}

type CartApplication interface {
	GetCart(cartRequest *cart_request.CartRequest) (cart_response.CartResponse, error)
}

func NewCartApplication(productRepository repository.ProductRepository) CartApplication {
	return &cartApplication{productRepository}
}

func (ca cartApplication) GetCart(cartRequest *cart_request.CartRequest) (cart_response.CartResponse, error) {
	products, err := ca.productRepository.FindAll()

	cartResponse := cart_response.CartResponse{
		Products: []product_response.ProductResponse{},
	}

	if err != nil {
		return cartResponse, err
	}

	for _, v := range products {
		for _, p := range cartRequest.Products {
			if p.ID == v.ID {
				productResponse := product_response.ProductResponse{
					ID:          v.ID,
					Quantity:    p.Quantity,
					UnityAmount: v.Amount,
					TotalAmount: v.Amount * p.Quantity,
					IsGift:      v.IsGift,
				}

				cartResponse.Products = append(cartResponse.Products, productResponse)
			}
		}
	}

	return cartResponse, nil
}
