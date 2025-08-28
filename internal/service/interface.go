package service

import "github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"

type ProductService interface {
	Add(product *model.ProductRequest) error
	GetAll() ([]model.Product, error)
	UpdateByID(ID uint, product model.ProductRequest) error
	DeleteByID(ID uint) error
}

type TransactionService interface {
	CreateTransaction(items []model.TransactionRequest) error
	GetHistory() ([]model.History, error)
}
