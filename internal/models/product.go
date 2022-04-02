package models

type Product struct {
	ID          int         `db:"id" json:"id"`
	Title       string      `db:"title" json:"title"`
	Description string      `db:"description" json:"description"`
	Price       int         `db:"price" json:"price"`
	Image       interface{} `db:"image" json:"image"`
}
