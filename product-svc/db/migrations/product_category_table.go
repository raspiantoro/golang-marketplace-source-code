package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/raspiantoro/golang-marketplace/product-svc/internal/model"
	"gorm.io/gorm"
)

func registerProductCategoryMigration() {
	migrations = append(migrations, CreateProductCategoryTable())
}

func CreateProductCategoryTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202408081642",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&model.ProductCategory{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("product-category")
		},
	}
}
