package mail

import (
	"github.com/tranvannghia021/gocore/src/service"
	"strings"
)

func RegisterMail(gmail string, signature string) error {
	mail := service.SMailer{
		To:      gmail,
		Cc:      nil,
		Subject: "Verify Mail In App",
		Body:    strings.Replace(templateRegisterHTML, "{{token}}", signature, 1),
	}
	return mail.Create().Send()
}
