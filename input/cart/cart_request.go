package cart

import (
	product_input "cart-checkout-simulation/input/product"
)

type CartRequest struct {
	Products []product_input.ProductRequest `json:"products"`
}
