package repository

import (
	"fmt"
	"log"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) Create(items []model.TransactionItem) error {
	// Mulai database transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		log.Printf("Error starting transaction: %v", tx.Error)
		return tx.Error
	}

	// Defer rollback jika ada panic atau error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to panic: %v", r)
		}
	}()

	// Buat transaksi baru
	transaction := model.Transaction{}
	result := tx.Create(&transaction)
	if result.Error != nil {
		tx.Rollback()
		log.Printf("Error creating transaction: %v", result.Error)
		if isDuplicateKey(result.Error) {
			return ErrDuplicate
		}
		return fmt.Errorf("%w: %v", ErrInternal, result.Error)
	}

	// 1. Kumpulkan semua product ID yang perlu di-lock
	var productIDs []uint
	for i, v := range items {
		productIDs = append(productIDs, v.ProductID)
		items[i].TransactionID = transaction.ID
	}

	// 2. Lock semua produk sekaligus di luar loop
	var products []model.Product

	result = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id IN ?", productIDs).
		Find(&products)

	if result.Error != nil {
		tx.Rollback()
		log.Printf("Error finding product: %v", result.Error)
		if isRecordNotFound(result.Error) {
			return ErrNotFound
		}
		return fmt.Errorf("%w: %v", ErrInternal, result.Error)
	}

	// 3. Loop untuk validasi stok
	for _, v := range items {
		for _, va := range products {
			if va.ID == v.ProductID {
				// Cek stock availability
				if va.Stock < v.Quantity {
					tx.Rollback()
					return fmt.Errorf("insufficient stock for %s", va.Name)
				}
				// Buat transaction_items
				if err := tx.Create(&v).Error; err != nil {
					tx.Rollback()
					log.Printf("Error creating transaction item: %v", err)
					if isDuplicateKey(err) {
						return ErrDuplicate
					}
					return fmt.Errorf("%w: %v", ErrInternal, err)
				}
				if err := tx.Model(&va).Update("stock", va.Stock-v.Quantity).Error; err != nil {
					tx.Rollback()
					if isDuplicateKey(err) {
						return ErrDuplicate
					}
					if isRecordNotFound(err) {
						return ErrNotFound
					}
					return fmt.Errorf("%w: %v", ErrInternal, err)
				}
			}
		}

	}

	// Commit transaction jika semua berhasil

	return tx.Commit().Error

}

func (r *transactionRepo) FindAll() ([]model.History, error) {
	var products []model.Product
	// var transactionItems []model.TransactionItem
	db := r.db

	db = db.Preload("TransactionItems")

	if err := db.Find(&products).Error; err != nil {
		log.Printf("Error finding product: %v", err)
		if isRecordNotFound(err) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	histories := []model.History{}
	for _, product := range products {
		// log.Println(product.Name)
		for _, item := range product.TransactionItems {
			histories = append(histories, model.History{
				Date:        item.CreatedAt,
				ProductName: product.Name,
				Quantity:    item.Quantity,
			})
		}
	}
	return histories, nil
}

// func (r *transactionRepo) FindAll() ([]model.Transaction, error) {
// 	var transactions []model.Transaction
// 	db := r.db

// 	db = db.Preload("Items").Preload("Items.Product")

// 	err := db.Find(&transactions).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return transactions, nil
// }

// func (r *transactionRepo) FindByID(id uint, preload bool) (*model.Transaction, error) {
// 	var transaction model.Transaction
// 	db := r.db

// 	if preload {
// 		db = db.Preload("Items").Preload("Items.Product")
// 	}

// 	err := db.First(&transaction, id).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &transaction, nil
// }

// func (r *transactionRepo) Updates(transaction *model.Transaction) error {
// 	return r.db.Save(transaction).Error
// }

// func (r *transactionRepo) Delete(id uint) error {
// 	return r.db.Transaction(func(tx *gorm.DB) error {
// 		// Delete items first (because of foreign key constraint)
// 		if err := tx.Where("transaction_id = ?", id).Delete(&model.TransactionItem{}).Error; err != nil {
// 			return err
// 		}

// 		// Then delete transaction header
// 		if err := tx.Delete(&model.Transaction{}, id).Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }
