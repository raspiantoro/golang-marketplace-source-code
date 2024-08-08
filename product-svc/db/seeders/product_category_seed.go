package seeders

import (
	"github.com/raspiantoro/golang-marketplace/product-svc/internal/model"
	"gorm.io/gorm"
)

func ProductCategorySeed(tx *gorm.DB) error {
	categories := []*model.ProductCategory{
		{
			Name: "Alat Tulis",
		},
		{
			Name: "Perlengkapan Rumah Tangga",
		},
		{
			Name: "Fashion",
		},
		{
			Name: "Elektronik",
		},
	}

	result := tx.Create(categories)

	return result.Error
}
