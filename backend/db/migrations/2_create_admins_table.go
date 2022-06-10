package migrations

import (
	"sister-backend/app/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateAdminsTable *gormigrate.Migration = &gormigrate.Migration{
	ID: "create_admins_table",
	Migrate: func(dbConn *gorm.DB) error {
		return dbConn.AutoMigrate(&entity.Admin{})
	},
	Rollback: func(dbConn *gorm.DB) error {
		return dbConn.Migrator().DropTable("admins")
	},
}
