package cart

import (
	product_request "cart-checkout-simulation/input/product/request"
)

type CartRequest struct {
	Products []product_request.ProductRequest `json:"products"`
}
