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
		TotalAmount:             0,
		TotalAmountWithDiscount: 0,
		TotalDiscount:           0,
		Products:                []product_response.ProductResponse{},
	}

	if err != nil {
		return cartResponse, err
	}

	for _, v := range products {
		for _, p := range cartRequest.Products {
			if p.ID == v.ID {
				var centsMath int32 = 100 * 100

				productDiscountInCents := int32(ca.discountService.GetDiscount(p.ID) * float32(centsMath))
				unitProductAmountInCents := v.Amount * centsMath
				totalProductAmountInCents := v.Amount * p.Quantity * centsMath

				productResponse := product_response.ProductResponse{
					ID:          v.ID,
					Quantity:    p.Quantity,
					UnityAmount: unitProductAmountInCents,
					TotalAmount: totalProductAmountInCents,
					Discount:    productDiscountInCents,
					IsGift:      v.IsGift,
				}

				cartResponse.TotalAmount += totalProductAmountInCents
				cartResponse.TotalDiscount += productDiscountInCents

				cartResponse.Products = append(cartResponse.Products, productResponse)
			}
		}
	}

	cartResponse.TotalAmountWithDiscount = cartResponse.TotalAmount - cartResponse.TotalDiscount

	return cartResponse, nil
}
