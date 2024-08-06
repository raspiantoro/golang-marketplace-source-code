package handler

import (
	"encoding/json"
	"net/http"

	"github.com/raspiantoro/golang-marketplace/product-svc/internal/contract"
	"github.com/raspiantoro/golang-marketplace/product-svc/internal/payload"
)

type ProductCategory struct {
	productCategoryServices contract.ProductCategoryService
}

func NewProductCategory(productCategoryServices contract.ProductCategoryService) *ProductCategory {
	return &ProductCategory{
		productCategoryServices: productCategoryServices,
	}
}

func (p *ProductCategory) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categories, err := p.productCategoryServices.GetAll(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Terjadi kesalahan pada server"))
		return
	}

	var responsePayloads []payload.GetProductCategoryResponse

	for _, category := range categories {
		responsePayloads = append(responsePayloads, payload.GetProductCategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	response, err := json.Marshal(&responsePayloads)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Terjadi kesalahan pada server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
