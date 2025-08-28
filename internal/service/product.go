package service

import (
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/repository"
	"gorm.io/gorm"
)

type productService struct {
	repo repository.ProductRepo
}

func NewProductService(repo repository.ProductRepo) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Add(product *model.ProductRequest) error {
	p := model.Product{
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
	if err := s.repo.Add(&p); err != nil {
		return err
	}
	return nil
}

func (s *productService) GetAll() ([]model.Product, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) UpdateByID(ID uint, product model.ProductRequest) error {
	prod := model.Product{
		Model: gorm.Model{ID: ID},
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
	if err := s.repo.Updates(&prod); err != nil {
		return err
	}
	return nil
}

func (s *productService) DeleteByID(ID uint) error {
	if err := s.repo.Delete(ID); err != nil {
		return err
	}
	return nil
}
