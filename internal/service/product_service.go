package service

import (
	"backend/internal/dto"
	"backend/internal/errors"
	"backend/internal/models"
	"backend/internal/repository"
	"log"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (ps *ProductService) GetProducts() ([]*models.Product, *errors.ApiError) {
	products, err := ps.ProductRepository.GetProducts()
	if err != nil {
		log.Println("unable to get users: " + err.Error())
		return nil, errors.InternalServerError(err)
	}

	return products, nil
}

func (ps *ProductService) AddProduct(productData *dto.ProductData) (int64, *errors.ApiError) {
	product := &models.Product{
		Title:       productData.Title,
		Description: productData.Description,
		Price:       productData.Price,
		Image:       productData.Image,
	}

	id, err := ps.ProductRepository.SaveProduct(product)
	if err != nil {
		log.Println("unable to save product: " + err.Error())
		return -1, errors.InternalServerError(err)
	}

	return id, nil
}
