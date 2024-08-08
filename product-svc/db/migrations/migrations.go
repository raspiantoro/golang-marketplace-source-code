package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

var migrations []*gormigrate.Migration

func setup() {
	registerProductCategoryMigration()
}

func Get() []*gormigrate.Migration {
	setup()
	return migrations
}
