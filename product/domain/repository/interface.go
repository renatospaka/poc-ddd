package repository

import "poc-ddd/product/domain/entity"

type ProductRepositoryInterface interface {
	Save(product *entity.Product)
	FindById(id string) (*entity.Product, error)
}
