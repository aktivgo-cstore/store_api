package repository

import (
	"github.com/jmoiron/sqlx"
	"store_api/internal/models"
)

type ProductRepository struct {
	MySqlConn *sqlx.DB
}

func NewProductRepository(mySqlConn *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		MySqlConn: mySqlConn,
	}
}

func (pr *ProductRepository) GetProducts() ([]*models.Product, error) {
	sql := `
		SELECT * FROM products
	`

	var products []*models.Product
	if err := pr.MySqlConn.Select(&products, sql); err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *ProductRepository) SaveProduct(product *models.Product) (int64, error) {
	sql := `
		INSERT INTO products
		(title, description, price, image) VALUE 
		(?, ?, ?, ?)
	`
	var args []interface{}
	args = append(args, product.Title, product.Description, product.Price, product.Image)

	result, err := pr.MySqlConn.Exec(sql, args...)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
