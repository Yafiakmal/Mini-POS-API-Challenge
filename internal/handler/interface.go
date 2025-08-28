package handler

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	CreateProduct(c *gin.Context)
	// GetProduct(c *gin.Context)
	GetAllProducts(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type TransactionHandler interface {
	CreateTransaction(c *gin.Context)
	GetTransactionHistory(c *gin.Context)
}
