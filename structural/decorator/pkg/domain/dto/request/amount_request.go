package request

type AmountRequest struct {
	Amount           float64 `json:"amount"`
	IsBrandDiscount  bool    `json:"isBrandDiscount"`
	IsBasketDiscount bool    `json:"isBasketDiscount"`
}
