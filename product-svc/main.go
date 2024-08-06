package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raspiantoro/golang-marketplace/product-svc/internal/handler"
	"github.com/raspiantoro/golang-marketplace/product-svc/internal/repository"
	"github.com/raspiantoro/golang-marketplace/product-svc/internal/service"
)

func main() {
	r := chi.NewRouter()

	productCategoryRepo := repository.NewProductCategory()
	productCategorySvc := service.NewProductCategory(productCategoryRepo)
	productCategoryHandler := handler.NewProductCategory(productCategorySvc)

	r.Get("/api/product-category", productCategoryHandler.GetAll)

	fmt.Println("your server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
