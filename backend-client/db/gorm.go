package db

import (
	"fmt"
	"log"
	"os"
	"sister-backend/config"
	"sister-backend/exception"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormConnection(conf config.Config) *gorm.DB {
	var dsn string

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level    // Ignore ErrRecordNotFound error for logger
			Colorful:      false,       // Disable color
		},
	)

	dbConfig := map[string]string{
		"host":     "POSTGRES_HOST",
		"port":     "POSTGRES_PORT",
		"user":     "POSTGRES_USER",
		"password": "POSTGRES_PASSWORD",
		"dbname":   "POSTGRES_DB",
		"sslmode":  "POSTGRES_SSLMODE",
		"TimeZone": "TIMEZONE",
	}

	for key, val := range dbConfig {
		env := conf.Get(val)
		exception.PanicWhenEnvNil(val, env)
		dsn += fmt.Sprintf("%s=%s ", key, env)
		dsn = strings.Trim(dsn, "")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	exception.PanicWhenError(err)

	return db
}
