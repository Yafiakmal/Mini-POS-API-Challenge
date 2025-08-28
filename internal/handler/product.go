package handler

import (
	"net/http"
	"strconv"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/service"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	s service.ProductService
}

func NewProductHandler(s service.ProductService) ProductHandler {
	return &productHandler{s: s}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var req model.ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add product
	if err := h.s.Add(&req); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

// func (h *productHandler) GetProduct(c *gin.Context) {
// 	idParam := c.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"id": uint(id)})
// }

func (h *productHandler) GetAllProducts(c *gin.Context) {
	products, err := h.s.GetAll()
	if err != nil {
		c.JSON(http.StatusFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var req model.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.s.UpdateByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Successfully"})
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.s.DeleteByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
