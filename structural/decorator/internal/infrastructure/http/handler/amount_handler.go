package handler

import (
	"design-patterns-go/structural/decorator/pkg/domain/dto/request"
	"design-patterns-go/structural/decorator/pkg/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CalculateDiscountHandler(ctx *gin.Context) {
	var req request.AmountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error json parse": err.Error()})
		return
	}

	discountService := service.NewDiscountService()
	amountResponse, err := discountService.CalculateDiscountedAmount(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, amountResponse)
}
