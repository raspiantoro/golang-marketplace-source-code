package repository

import (
	"context"

	"github.com/raspiantoro/golang-marketplace/product-svc/internal/model"
)

type ProductCategory struct{}

func NewProductCategory() *ProductCategory {
	return &ProductCategory{}
}

func (p *ProductCategory) GetAll(ctx context.Context) ([]model.ProductCategory, error) {
	categories := []model.ProductCategory{
		{
			ID:   1,
			Name: "Alat Tulis",
		},
		{
			ID:   2,
			Name: "Perlengkapan Rumah Tangga",
		},
		{
			ID:   3,
			Name: "Fashion",
		},
	}

	return categories, nil
}
