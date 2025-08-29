package service

import (
	"errors"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/repository"
)

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) CreateTransaction(items []model.TransactionRequest) error {
	transactionItem := []model.TransactionItem{}
	for _, v := range items {
		transactionItem = append(transactionItem, model.TransactionItem{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}
	if err := s.repo.Create(transactionItem); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		} else if errors.Is(err, repository.ErrDuplicate) {
			return ErrDuplicate
		}
		return ErrInternal
	}

	return nil
}

func (s *transactionService) GetHistory() ([]model.History, error) {
	histories, err := s.repo.FindAll()
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
	}
	return histories, nil
}
