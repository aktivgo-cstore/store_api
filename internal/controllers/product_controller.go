package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"store_api/internal/dto"
	"store_api/internal/helpers"
	"store_api/internal/service"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (pc *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, er := service.GetToken(r.Header)
	if er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	if er = service.CheckToken(token); er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	products, er := pc.ProductService.GetProducts()
	if er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	encode, err := json.Marshal(products)
	if err != nil {
		log.Println("unable to encode products: " + err.Error())
		helpers.ErrorResponse(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(encode)
}

func (pc *ProductController) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token, er := service.GetToken(r.Header)
	if er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	if er = service.CheckAccess(token); er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("unable to read request body: " + err.Error())
		helpers.ErrorResponse(w, "Некорректный запрос", http.StatusInternalServerError)
		return
	}

	var productData *dto.ProductData
	if err = json.Unmarshal(body, &productData); err != nil {
		log.Println("unable to decode request body: " + err.Error())
		helpers.ErrorResponse(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	id, er := pc.ProductService.AddProduct(productData)
	if er != nil {
		helpers.ErrorResponse(w, er.Message, er.Status)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"id": %d}`, id)))
}
