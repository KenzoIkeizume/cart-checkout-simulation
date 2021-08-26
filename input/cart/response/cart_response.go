package cart

import (
	product_response "cart-checkout-simulation/input/product/response"
)

type CartResponse struct {
	Products []product_response.ProductResponse `json:"products"`
}
