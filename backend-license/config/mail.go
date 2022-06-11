package config

import (
	"strconv"
)

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewMailConfig(conf Config) (mailConfig MailConfig) {
	mailConfig.Host = conf.Get("MAIL_HOST")
	mailConfig.Port, _ = strconv.Atoi(conf.Get("MAIL_PORT"))
	mailConfig.Username = conf.Get("MAIL_USERNAME")
	mailConfig.Password = conf.Get("MAIL_PASSWORD")

	return mailConfig
}
