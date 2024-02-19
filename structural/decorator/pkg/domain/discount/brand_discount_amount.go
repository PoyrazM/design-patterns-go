package discount

import (
	"design-patterns-go/structural/decorator/pkg/domain/dto/response"
	"design-patterns-go/structural/decorator/pkg/domain/model"
	"design-patterns-go/structural/decorator/ports"
)

const brandDiscountPercentage float32 = 0.25

type BrandDiscountAmountDecorator struct {
	amount ports.IAmount
}

func NewBrandDiscountDecorator(a ports.IAmount) *BrandDiscountAmountDecorator {
	return &BrandDiscountAmountDecorator{amount: a}
}

func (bda *BrandDiscountAmountDecorator) GetAmount() *response.AmountResponse {
	previousAmount := bda.amount.GetAmount()
	amount := model.NewAmount(previousAmount.Amount)
	discountedAmount := amount.CalculateDiscountedAmount(brandDiscountPercentage)
	return &response.AmountResponse{
		Amount:                  discountedAmount,
		TotalDiscountPercentage: previousAmount.TotalDiscountPercentage + brandDiscountPercentage,
		Description:             previousAmount.Description + " | Added the brand discount.",
	}
}
