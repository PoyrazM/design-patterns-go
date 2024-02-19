package model

import (
	"math"
)

type Amount struct {
	Value float64
}

func NewAmount(amount float64) *Amount {
	return &Amount{Value: amount}
}

func (a *Amount) CalculateDiscountedAmount(discountPercentage float32) float64 {
	discountedAmount := a.Value - (a.Value * float64(discountPercentage))
	roundedAmount := math.Round(discountedAmount*100) / 100
	return roundedAmount
}
