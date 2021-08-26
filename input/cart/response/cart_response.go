package cart

import (
	product_response "cart-checkout-simulation/input/product/response"
)

type CartResponse struct {
	TotalAmount             int32                              `json:"total_amount"`
	TotalAmountWithDiscount int32                              `json:"total_amount_with_discount"`
	TotalDiscount           int32                              `json:"total_discount"`
	Products                []product_response.ProductResponse `json:"products"`
}
