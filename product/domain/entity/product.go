package entity

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID          string
	Description string
	Price       float32
	MaxDiscount float32
	valid       bool
}

func NewProduct() *Product {
	fmt.Println("NewProduct Created")
	return &Product{
		ID:          uuid.NewV4().String(),
		Description: "",
		Price:       0,
		MaxDiscount: 0,
		valid:       false,
	}
}

func (p *Product) IsValid() bool {
	return true
}
