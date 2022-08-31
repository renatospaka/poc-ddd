package postgres

import (
	"database/sql"
	"fmt"
	"poc-ddd/product/domain/entity"
)

type ProductRepositoryDb struct {
	db *sql.DB
}

func NewProductRepositoryDb(db *sql.DB) *ProductRepositoryDb {
	fmt.Println("NewProductRepositoryDb Created")
	return &ProductRepositoryDb{db: db}
}

func (r *ProductRepositoryDb) Save(product *entity.Product) {
	fmt.Printf("Product %v Saved\n", product)
}

func (r *ProductRepositoryDb) FindById(id string) (*entity.Product, error) {
	fmt.Printf("Product %s Found\n", id)
	product := entity.NewProduct()
	return product, nil
}
