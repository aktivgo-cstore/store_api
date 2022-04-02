package app

import (
	"backend/internal/controllers"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/internal/storage/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	port         = os.Getenv("API_PORT")
	mySqlConnStr = os.Getenv("MYSQL_CONN_STR")
)

func Run() error {
	router := mux.NewRouter()
	mySqlConn, err := mysql.CreateConnection(mySqlConnStr)
	if err != nil {
		return err
	}

	productRepository := repository.NewProductRepository(mySqlConn)
	productService := service.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	router.HandleFunc("/products", productController.GetProducts).Methods("GET")
	router.HandleFunc("/products", productController.AddProduct).Methods("POST")

	log.Println("Store api server started on port " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		return err
	}

	return nil
}
