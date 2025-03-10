package routes

import (
	"github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/http/v1/rest/gin/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, transactionController *controller.TransactionController) {
	v1 := router.Group("/v1")
	{
		v1.GET("/transactions", transactionController.GetTransactions)
		v1.POST("/transactions", transactionController.CreateTransaction)
	}
}
