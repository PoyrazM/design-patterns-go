package service

import (
	"design-patterns-go/structural/decorator/pkg/domain/discount"
	"design-patterns-go/structural/decorator/pkg/domain/dto/request"
	"design-patterns-go/structural/decorator/pkg/domain/dto/response"
	"design-patterns-go/structural/decorator/ports"
)

type DiscountService struct {
}

func NewDiscountService() *DiscountService {
	return &DiscountService{}
}

func (ds *DiscountService) CalculateDiscountedAmount(req *request.AmountRequest) (*response.AmountResponse, error) {
	var amount ports.IAmount
	amount = discount.NewBaseAmount(req)

	if req.IsBrandDiscount {
		amount = discount.NewBrandDiscountDecorator(amount)
	}

	if req.IsBasketDiscount {
		amount = discount.NewBasketDiscountDecorator(amount)
	}

	return amount.GetAmount(), nil
}
