package service

import (
	"context"

	"github.com/raspiantoro/golang-marketplace/product-svc/internal/contract"
	"github.com/raspiantoro/golang-marketplace/product-svc/internal/model"
)

type ProductCategory struct {
	productCategoryRepo contract.ProductCategoryRepository
}

func NewProductCategory(productCategoryRepo contract.ProductCategoryRepository) *ProductCategory {
	return &ProductCategory{
		productCategoryRepo: productCategoryRepo,
	}
}

func (p *ProductCategory) GetAll(ctx context.Context) ([]model.ProductCategory, error) {
	return p.productCategoryRepo.GetAll(ctx)
}
