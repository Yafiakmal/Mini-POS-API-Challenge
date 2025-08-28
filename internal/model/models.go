package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(100);not null" json:"name"`
	Price float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock uint    `gorm:"not null" json:"stock"`

	// Relationship dengan TransactionItem
	TransactionItems []TransactionItem `gorm:"foreignKey:ProductID"`
}

// Model Transaction
type Transaction struct {
	gorm.Model

	// Relationship dengan TransactionItem
	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionID"`
}

// Model TransactionItem (Pivot Table)
type TransactionItem struct {
	gorm.Model
	TransactionID uint `gorm:"not null;index"`
	ProductID     uint `gorm:"not null;index"`
	Quantity      uint `gorm:"not null"`

	// Foreign Key Relationships
	Transaction Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	Product     Product     `gorm:"foreignKey:ProductID;references:ID"`
}

type History struct {
	Date        time.Time
	ProductName string
	Quantity    uint
}

type ProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Stock uint    `json:"stock" binding:"required"`
}

type TransactionRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}

// type Product struct {
// 	gorm.Model
// 	Name  string  `gorm:"type:varchar(100);not null;index;unique" json:"name"`
// 	Price float64 `gorm:"type:decimal(10,2);not null" json:"price"`
// 	Stock uint    `gorm:"not null" json:"stock"`
// }

// type Transaction struct {
// 	ID    uint              `gorm:"primaryKey" json:"id"`
// 	Date  time.Time         `gorm:"not null" json:"date"`
// 	Total float64           `json:"total"`
// 	Items []TransactionItem `gorm:"foreignKey:TransactionID" json:"items,omitempty"`
// }

// type TransactionItem struct {
// 	ID            uint    `gorm:"primaryKey" json:"id"`
// 	TransactionID uint    `gorm:"not null;index" json:"transaction_id"`
// 	ProductID     uint    `gorm:"not null;index" json:"product_id"`
// 	Quantity      int     `gorm:"not null;default:1" json:"quantity"`
// 	UnitPrice     float64 `gorm:"not null" json:"unit_price"`
// 	Product       Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
// }
