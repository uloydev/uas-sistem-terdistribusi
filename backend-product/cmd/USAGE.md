# Command Usage
## migration
### *migrate all migrations*
    go run cmd/migration.go up
### *rollback all migrations*
    go run cmd/migration.go down
### *refresh all migrations*
    go run cmd/migration.go refresh
## swagger
### update generated swagger specs (dev only)
    go run cmd/swag/swag.go init
