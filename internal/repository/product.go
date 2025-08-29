package repository

import (
	"fmt"
	"log"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) Add(product *model.Product) error {
	result := r.db.Create(product)

	if result.Error != nil {
		if isDuplicateKey(result.Error) {
			return ErrDuplicate
		}
		return fmt.Errorf("%w: %v", ErrInternal, result.Error)
	}
	return nil
}

func (r *productRepo) FindByID(id uint) (*model.Product, error) {
	product := model.Product{}
	res := r.db.First(&product, id)

	if res.Error != nil {
		if isRecordNotFound(res.Error) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrInternal, res.Error)
	}
	return &product, nil
}
func (r *productRepo) FindAll() ([]model.Product, error) {
	product := []model.Product{}
	res := r.db.Find(&product)
	if res.Error != nil {
		if isRecordNotFound(res.Error) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrInternal, res.Error)
	}
	return product, nil
}
func (r *productRepo) Updates(product *model.Product) error {
	res := r.db.Model(product).Updates(*product)
	if res.Error != nil {
		if isDuplicateKey(res.Error) {
			return ErrDuplicate
		}
		if isRecordNotFound(res.Error) {
			return ErrNotFound
		}
		return fmt.Errorf("%w: %v", ErrInternal, res.Error)
	}
	return nil
}
func (r *productRepo) Delete(id uint) error {
	log.Println("deleting :", id)
	res := r.db.Delete(&model.Product{}, id)
	if res.Error != nil {
		if isRecordNotFound(res.Error) {
			return ErrNotFound
		}
		return fmt.Errorf("%w: %v", ErrInternal, res.Error)
	}
	return nil
}
