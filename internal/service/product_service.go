package service

import "backend/internal/repository"

type ProductService struct {
	UserRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{
		UserRepository: productRepository,
	}
}
