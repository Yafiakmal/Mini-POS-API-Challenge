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
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}
func (r *productRepo) FindByID(id uint) (*model.Product, error) {
	product := model.Product{}
	res := r.db.First(&product, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &product, nil
}
func (r *productRepo) FindAll() ([]model.Product, error) {
	product := []model.Product{}
	if err := r.db.Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (r *productRepo) Updates(product *model.Product) error {
	if err := r.db.Model(product).Updates(*product).Error; err != nil {
		return err
	}
	return nil
}
func (r *productRepo) Delete(id uint) error {
	log.Println("deleting :", id)
	res := r.db.Delete(&model.Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected < 1 {
		return fmt.Errorf("product not found")
	}
	return nil
}
