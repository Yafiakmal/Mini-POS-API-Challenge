package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/handler"
)

func SetupProductRoutes(router *gin.Engine, handler handler.ProductHandler) {
	router.POST("/product", handler.CreateProduct)
	// router.GET("/product/:id", handler.GetProduct)
	router.GET("/products", handler.GetAllProducts)
	router.PUT("/product/:id", handler.UpdateProduct)
	router.DELETE("/product/:id", handler.DeleteProduct)
}
