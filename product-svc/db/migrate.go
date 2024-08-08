package main

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/raspiantoro/golang-marketplace/product-svc/db/migrations"
	"github.com/raspiantoro/golang-marketplace/product-svc/db/seeders"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=localhost user=admin password=admin dbname=product port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations.Get())

	if err = m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	s := seeders.NewSeeders(db, seeders.WithDefaultSeeders())

	if err = s.Seed(); err != nil {
		log.Fatalf("Seeders failed: %v", err)
	}

	log.Println("Migration and Seeders did run successfully")
}
