package model

import (
	"log"
	"os"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if os.Getenv("MODE") == "development" {
		err := db.Migrator().DropTable(&Product{}, &SalesTransaction{}, &TransactionItem{})
		if err != nil {
			log.Fatalf("failed to drop table: %v", err)
		}
	}

	return db.AutoMigrate(
		&Product{},
		&SalesTransaction{},
		&TransactionItem{},
		// tambahkan semua model di sini
	)
}
