package response

type AmountResponse struct {
	Amount                  float64 `json:"amount"`
	TotalDiscountPercentage float32 `json:"totalDiscountPercentage"`
	Description             string  `json:"description"`
}
