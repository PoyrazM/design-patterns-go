package router

import (
	"design-patterns-go/structural/decorator/internal/infrastructure/http/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func StartDiscountApp() {
	r := gin.Default()

	r.POST("/calculate-amount", handler.CalculateDiscountHandler)

	log.Fatal(r.Run(":9090"))
}
