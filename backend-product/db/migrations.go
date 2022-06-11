package db

import (
	"sister-backend/db/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations = []*gormigrate.Migration{

	migrations.CreateProductsTable,
}
