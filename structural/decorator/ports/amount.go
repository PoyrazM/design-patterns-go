package ports

import (
	"design-patterns-go/structural/decorator/pkg/domain/dto/response"
)

type IAmount interface {
	GetAmount() *response.AmountResponse
}
