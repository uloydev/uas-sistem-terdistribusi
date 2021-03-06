package migrations

import (
	"sister-backend/app/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateProductsTable *gormigrate.Migration = &gormigrate.Migration{
	ID: "create_products_table",
	Migrate: func(dbConn *gorm.DB) error {
		return dbConn.AutoMigrate(&entity.Product{})
	},
	Rollback: func(dbConn *gorm.DB) error {
		return dbConn.Migrator().DropTable("products")
	},
}
