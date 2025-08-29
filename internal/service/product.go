package service

import (
	"errors"

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
		switch {
		case errors.Is(err, repository.ErrDuplicate):
			return ErrDuplicate
		default:
			return ErrInternal
		}
	}
	return nil
}

func (s *productService) GetAll() ([]model.Product, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrInternal
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
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		} else if errors.Is(err, repository.ErrDuplicate) {
			return ErrDuplicate
		}
		return ErrInternal
	}
	return nil
}

func (s *productService) DeleteByID(ID uint) error {
	if err := s.repo.Delete(ID); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		} else if errors.Is(err, repository.ErrDuplicate) {
			return ErrDuplicate
		}
		return ErrInternal
	}
	return nil
}
