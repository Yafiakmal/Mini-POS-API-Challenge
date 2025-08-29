package model

import (
	"log"
	"os"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if os.Getenv("APP_ENV") == "development" {
		err := db.Migrator().DropTable(&Product{}, &Transaction{}, &TransactionItem{})
		if err != nil {
			log.Fatalf("failed to drop table: %v", err)
		}
	}

	if err := db.AutoMigrate(
		&Product{},
		&Transaction{},
		&TransactionItem{},
		// tambahkan semua model di sini
	); err != nil {
		return err
	}
	db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_products_name_unique ON products (name) WHERE deleted_at IS NULL;")
	return nil
}
