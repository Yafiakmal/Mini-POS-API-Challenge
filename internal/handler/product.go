package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/service"
	"github.com/yafiakmal/Mini-POS-API-Challenge/util"

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
		util.Error(c, http.StatusBadRequest, "body request not valid", err)
		return
	}

	// Add product
	if err := h.s.Add(&req); err != nil {
		if errors.Is(err, service.ErrDuplicate) {
			util.Error(c, http.StatusConflict, "cannot assign duplicate product", err)
			return
		}
		util.Error(c, http.StatusInternalServerError, "internal server error", err)
		return
	}
	util.Created(c, req)
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
		if errors.Is(err, service.ErrNotFound) {
			util.Error(c, http.StatusNotFound, "not found", err)
			return
		}
		util.Error(c, http.StatusInternalServerError, "error", err)
		return
	}

	util.Success(c, products)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "invalid id", err)
		return
	}
	var req model.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, http.StatusBadRequest, "body request not valid", err)
		return
	}

	if err := h.s.UpdateByID(uint(id), req); err != nil {
		if errors.Is(err, service.ErrDuplicate) {
			util.Error(c, http.StatusConflict, "cannot assign duplicate product", err)
			return
		}
		util.Error(c, http.StatusInternalServerError, "error", err)
		return
	}
	util.Success(c, req)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		util.Error(c, http.StatusBadRequest, "invalid id", err)
		return
	}

	if err := h.s.DeleteByID(uint(id)); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			util.Error(c, http.StatusNotFound, "not found", err)
			return
		}
		util.Error(c, http.StatusInternalServerError, "error", err)
		return
	}
	util.Success(c, nil)
}
