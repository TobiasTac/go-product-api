package main

import (
	"net/http"

	"github.com/TobiasTac/go-product-api/configs"
	"github.com/TobiasTac/go-product-api/internal/entity"
	"github.com/TobiasTac/go-product-api/internal/infra/database"
	"github.com/TobiasTac/go-product-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)

	http.ListenAndServe(":8000", r)
}
