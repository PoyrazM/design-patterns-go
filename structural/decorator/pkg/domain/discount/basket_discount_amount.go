package discount

import (
	"design-patterns-go/structural/decorator/pkg/domain/dto/response"
	"design-patterns-go/structural/decorator/pkg/domain/model"
	"design-patterns-go/structural/decorator/ports"
)

const basketDiscountPercentage float32 = 0.10

type BasketDiscountAmountDecorator struct {
	amount ports.IAmount
}

func NewBasketDiscountDecorator(a ports.IAmount) *BasketDiscountAmountDecorator {
	return &BasketDiscountAmountDecorator{amount: a}
}

func (bda *BasketDiscountAmountDecorator) GetAmount() *response.AmountResponse {
	previousAmount := bda.amount.GetAmount()
	amount := model.NewAmount(previousAmount.Amount)
	discountedAmount := amount.CalculateDiscountedAmount(basketDiscountPercentage)
	return &response.AmountResponse{
		Amount:                  discountedAmount,
		TotalDiscountPercentage: previousAmount.TotalDiscountPercentage + basketDiscountPercentage,
		Description:             previousAmount.Description + " | Added the basket discount.",
	}
}
