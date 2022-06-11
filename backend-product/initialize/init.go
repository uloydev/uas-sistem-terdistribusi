package initialize

import (
	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RunInitFunctions(app *fiber.Group, DB *gorm.DB, MailConn *gomail.Dialer) {
	for _, initFunction := range InitFunctions {
		initFunction(app, DB, MailConn)
	}
}
