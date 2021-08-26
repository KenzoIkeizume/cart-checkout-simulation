package product

type ProductResponse struct {
	ID          int32 `json:"id"`
	Quantity    int32 `json:"quantity"`
	UnityAmount int32 `json:"unit_amount"`
	TotalAmount int32 `json:"total_amount"`
	Discount    int32 `json:"discount"`
	IsGift      bool  `json:"is_gift"`
}
