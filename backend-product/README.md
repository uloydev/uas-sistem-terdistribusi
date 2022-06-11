# sister-backend
>sister-backend is template for built web api with golang programming language to aim best performance and speed of development. Feel free to use this boilerplate for your project
## used package
    - gofiber
    - gorm
    - gormigrate
    - golang-jwt
    - swaggo
    - ozzo-validation
    - gomail
    - godotenv
## How To Run

### copy .env.example to .env and edit your .env config
    cp .env.example .env
### install dependencies
    go mod tidy
### run database migration (if needed) more details at [commands](cmd/USAGE.md)
    go run cmd/migration.go up
### run server
    go run main.go
### go to api docs (open in browser)
    http://localhost:6991/swagger