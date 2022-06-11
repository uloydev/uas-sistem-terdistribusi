package migrations

import (
	"sister-backend/app/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateLicensesTable *gormigrate.Migration = &gormigrate.Migration{
	ID: "create_licenses_table",
	Migrate: func(dbConn *gorm.DB) error {
		return dbConn.AutoMigrate(&entity.License{})
	},
	Rollback: func(dbConn *gorm.DB) error {
		return dbConn.Migrator().DropTable("licenses")
	},
}
