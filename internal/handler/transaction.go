package handler

import (
	"net/http"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/service"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	s service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) TransactionHandler {
	return &transactionHandler{s: s}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var req []model.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.s.CreateTransaction(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (h *transactionHandler) GetTransactionHistory(c *gin.Context) {
	histories, err := h.s.GetHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"histories": histories})
}
