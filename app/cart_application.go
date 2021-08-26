package app

import (
	"time"

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

	winGift := false

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
					IsGift:      false,
				}

				cartResponse.TotalAmount += totalProductAmountInCents
				cartResponse.TotalDiscount += productDiscountInCents

				cartResponse.Products = append(cartResponse.Products, productResponse)
			}
		}

		if ca.isBlackFriday() && v.IsGift && !winGift {
			productResponse := product_response.ProductResponse{
				ID:          v.ID,
				Quantity:    1,
				UnityAmount: 0,
				TotalAmount: 0,
				Discount:    0,
				IsGift:      false,
			}

			cartResponse.Products = append(cartResponse.Products, productResponse)
			winGift = true

			continue
		}
	}

	cartResponse.TotalAmountWithDiscount = cartResponse.TotalAmount - cartResponse.TotalDiscount

	return cartResponse, nil
}

func (ca cartApplication) isBlackFriday() bool {
	month := 9
	day := 26

	_, m, d := time.Now().Date()

	return int(m) == month && d == day
}
