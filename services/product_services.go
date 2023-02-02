package services

import (
	"errors"
	"we_bid/mock"
	"we_bid/models"
)

type ProductService interface {
	GetProductsList(f ProductsFilter) (*[]models.Product, error)
	GetProductByID(id int) (*models.Product, error)
}

type productService struct {
}

func CreateProductService() ProductService {
	return &productService{}
}

type ProductsFilter struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (s *productService) GetProductsList(f ProductsFilter) (*[]models.Product, error) {
	return &mock.ProductList, nil
}

func (s *productService) GetProductByID(id int) (*models.Product, error) {
	var result *models.Product = nil
	for _, p := range mock.ProductList {
		if p.ID == id {
			result = &p
		}
	}
	if result == nil {
		return nil, errors.New("product not found")
	}
	return &mock.ProductList[0], nil
}
