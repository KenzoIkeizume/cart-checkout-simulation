package app

import (
	discount_service "cart-checkout-simulation/infra/discount"
	repository "cart-checkout-simulation/infra/repository"
	cart_request "cart-checkout-simulation/input/cart/request"
	cart_response "cart-checkout-simulation/input/cart/response"
	product_response "cart-checkout-simulation/input/product/response"
)

type cartApplication struct {
	productRepository repository.ProductRepository
	discountService   discount_service.DiscountService
}

type CartApplication interface {
	GetCart(cartRequest *cart_request.CartRequest) (cart_response.CartResponse, error)
}

func NewCartApplication(productRepository repository.ProductRepository, discountService discount_service.DiscountService) CartApplication {
	return &cartApplication{productRepository, discountService}
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
				var centsMath int32 = 100 * 100
				productDiscount := int32(ca.discountService.GetDiscount(p.ID) * float32(centsMath))

				productResponse := product_response.ProductResponse{
					ID:          v.ID,
					Quantity:    p.Quantity,
					UnityAmount: v.Amount * centsMath,
					TotalAmount: v.Amount * p.Quantity * centsMath,
					Discount:    productDiscount,
					IsGift:      v.IsGift,
				}

				cartResponse.Products = append(cartResponse.Products, productResponse)
			}
		}
	}

	return cartResponse, nil
}
