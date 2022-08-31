package main

import (
	"database/sql"
	"fmt"
	"poc-ddd/product/domain/entity"
	// "poc-ddd/product/domain/service"
	"poc-ddd/product/infrastructure/repository/postgres"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "local_dev"
	password = "thisislocal#noneedtoworry"
	dbname   = "miggos"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic("deu ruim na conexão com o banco" + err.Error())
	}
	defer db.Close()
	fmt.Println("conexão com o banco estabelecida com sucesso")

	prodRepository := postgres.NewProductRepositoryDb(db)
	// prodService := service.NewProductService(prodRepository)
	product := entity.NewProduct()
	product.Description = "Um produto"
	product.Price = 344
	product.MaxDiscount = 12.5

	// prodService.SaveProduct(product)
	prodRepository.Save(product)

	db.Close()
}