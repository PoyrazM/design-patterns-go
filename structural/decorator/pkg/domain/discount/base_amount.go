package discount

import (
	"design-patterns-go/structural/decorator/pkg/domain/dto/request"
	"design-patterns-go/structural/decorator/pkg/domain/dto/response"
	"design-patterns-go/structural/decorator/pkg/domain/model"
)

const baseDiscountPercentage float32 = 0.0

type BaseAmount struct {
	request *request.AmountRequest
}

func NewBaseAmount(r *request.AmountRequest) *BaseAmount {
	return &BaseAmount{request: r}
}

func (ba *BaseAmount) GetAmount() *response.AmountResponse {
	amount := model.NewAmount(ba.request.Amount)
	discountedAmount := amount.CalculateDiscountedAmount(baseDiscountPercentage)
	return &response.AmountResponse{
		Amount:                  discountedAmount,
		TotalDiscountPercentage: baseDiscountPercentage,
		Description:             "Calculated the amount with base discount percentage.",
	}
}
