package main

import (
	"fmt"
	"os"
	"sister-backend/config"
	"sister-backend/db"

	"github.com/go-gormigrate/gormigrate/v2"
)

func migrateRun(migrator *gormigrate.Gormigrate) {
	if err := migrator.Migrate(); err != nil {
		panic(err)
	}
	fmt.Println("SUCCESS migrate migrations")
}

func rollbackRun(migrator *gormigrate.Gormigrate) {
	for _, migration := range db.Migrations {
		if err := migrator.RollbackMigration(migration); err != nil {
			panic(err)
		}
	}
	fmt.Println("SUCCESS rollback migrations")
}

func main() {

	args := os.Args

	if len(args) < 2 {
		panic("not enough parameter")
	}

	conf := config.New()

	dbConn := db.NewGormConnection(conf)

	migrator := gormigrate.New(dbConn, gormigrate.DefaultOptions, db.Migrations)

	switch args[1] {
	case "up":
		migrateRun(migrator)
	case "down":
		rollbackRun(migrator)
	case "refresh":
		rollbackRun(migrator)
		migrateRun(migrator)
	default:
		panic("invalid command, example: up | down | refresh")
	}

}
