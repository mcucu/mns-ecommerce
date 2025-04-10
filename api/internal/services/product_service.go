package services

import (
	"errors"

	"github.com/mcucu/mns-ecommerce/internal/models"
)

type IProductService interface {
	GetAll() ([]models.Product, error)
}

type ProductService struct {
	Options
}

func NewProductService(opt Options) IProductService {
	return &ProductService{
		Options: opt,
	}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	products, err := s.Product.GetAll()
	if err != nil {
		return nil, errors.New("failed to fetch products")
	}

	return products, nil
}
