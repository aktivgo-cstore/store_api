package controllers

import (
	"backend/internal/service"
	"net/http"
)

type ProductController struct {
	UserService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{
		UserService: productService,
	}
}

func (pc *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}
