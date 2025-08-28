package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/handler"
)

func SetupTransactionRoutes(router *gin.Engine, handler handler.TransactionHandler) {
	router.POST("/transaction", handler.CreateTransaction)
	router.GET("/transactions", handler.GetTransactionHistory)
}
