package repository

import (
	"time"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
	_ "gorm.io/gorm"
)

type Transaction struct {
	Time time.Time
}

type ProductRepo interface {
	Add(product *model.Product) error
	FindByID(id uint) (*model.Product, error)
	FindAll() ([]model.Product, error)
	Updates(product *model.Product) error
	Delete(id uint) error
}

type TransactionRepository interface {
	Create(items []model.TransactionItem) error
	// FindByID(id uint, preload bool) (*model.Transaction, error)
	FindAll() ([]model.History, error)
	// Updates(transaction *model.Transaction) error
	// Delete(id uint) error
}
