package repository

import "github.com/jmoiron/sqlx"

type ProductRepository struct {
	MySqlConn *sqlx.DB
}

func NewProductRepository(mySqlConn *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		MySqlConn: mySqlConn,
	}
}
