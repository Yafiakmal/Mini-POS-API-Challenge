package handler

import (
	"errors"
	"net/http"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/service"
	"github.com/yafiakmal/Mini-POS-API-Challenge/util"

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
		util.Error(c, http.StatusBadRequest, "invalid id", err)
		return
	}

	if err := h.s.CreateTransaction(req); err != nil {
		if errors.Is(err, service.ErrDuplicate) {
			util.Error(c, http.StatusConflict, "cannot assign duplicate product", err)
			return
		}
		util.Error(c, http.StatusInternalServerError, "error", err)
		return
	}

	util.Created(c, req)
}

func (h *transactionHandler) GetTransactionHistory(c *gin.Context) {
	histories, err := h.s.GetHistory()
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			util.Error(c, http.StatusNotFound, "not found", err)
			return
		}
		util.Error(c, http.StatusInternalServerError, "error", err)
		return
	}

	util.Success(c, histories)
}
