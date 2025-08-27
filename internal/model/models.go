package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(100);not null;index;unique" json:"name"`
	Price float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock uint    `gorm:"not null" json:"stock"`
}

type SalesTransaction struct {
	gorm.Model
}

type TransactionItem struct {
	Amount    uint `gorm:"not null" json:"amount"`
	STID      uint `gorm:"primaryKey;not null" json:"stid"`
	ProductID uint `gorm:"primaryKey;not null" json:"product_id"`

	// relation
	SalesTransaction SalesTransaction `gorm:"foreignKey:STID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Product          Product          `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
