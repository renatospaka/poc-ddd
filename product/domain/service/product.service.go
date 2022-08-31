package service

import (
	"fmt"
	"poc-ddd/product/domain/entity"
	"poc-ddd/product/domain/repository"
)

type ProductService struct {
	repo repository.ProductRepositoryInterface
}

func NewProductService(repo *repository.ProductRepositoryInterface) *ProductService{
	fmt.Println("NewProductService Created")
	return &ProductService{repo: *repo}
}

func (s *ProductService) SaveProduct(product *entity.Product) bool {
	return true
}

func (s *ProductService) DeleteProduct(id string) bool {
	return true
}