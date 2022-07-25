package inject

import (
	"gopkg.in/gomail.v2"
)

func GetMail() *gomail.Dialer {
	if service.mail == nil {
		cfg := GetConfig()
		service.mail = gomail.NewDialer(cfg.Mail.Host, cfg.Mail.Port, cfg.Mail.User, cfg.Mail.Pwd)
	}
	return service.mail
}
