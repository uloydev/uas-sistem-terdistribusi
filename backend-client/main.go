package main

import (
	"sister-backend/config"
	"sister-backend/db"
	"sister-backend/initialize"
	"sister-backend/job"
	"sister-backend/utils"
	"time"

	_ "sister-backend/docs"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title                       SISTEM TERDISTRIBUSI API
// @version                     1.0
// @description                 SISTEM TERDISTRIBUSI API BACKEND.
// @contact.name                uloydev
// @contact.email               uloydev.gmail.com
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath                    /
func main() {
	conf := config.New()
	dbConn := db.NewGormConnection(conf)
	mailConf := config.NewMailConfig(conf)
	mailConn := utils.NewMailConnection(mailConf)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTION",
	}))

	v1 := app.Group("/v1")

	initialize.RunInitFunctions(v1.(*fiber.Group), dbConn, mailConn)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:          "http://localhost:8693/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	// register jobs
	s := gocron.NewScheduler(time.UTC)
	s.Every(3).Hour().Do(job.SyncProduct)
	s.StartAsync()

	app.Listen(":8693")
}
