package contract

import (
	"context"

	"github.com/raspiantoro/golang-marketplace/product-svc/internal/model"
)

type ProductCategoryRepository interface {
	GetAll(ctx context.Context) ([]model.ProductCategory, error)
}
