package service

import (
	"github.com/tranvannghia021/gocore/vars"
	mail "github.com/xhit/go-simple-mail/v2"
	"os"
)

type Imailer interface {
	Create(sMailer *SMailer) *mail.Email
	Send() error
}

type SMailer struct {
	mailClient *mail.SMTPClient
	From       string
	To         string
	Cc         []string
	Subject    string
	Body       string
	Attachment string
	email      *mail.Email
}

func (s *SMailer) Create() *SMailer {
	s.mailClient = vars.Mail
	email := mail.NewMSG()
	if s.From == "" {
		s.From, _ = os.LookupEnv("MAIL_USERNAME")
	}
	email.SetFrom(s.From)
	email.AddTo(s.To)
	for _, v := range s.Cc {
		email.AddCc(v)
	}
	email.SetSubject(s.Subject)
	email.SetBody(mail.TextHTML, s.Body)
	if s.Attachment != "" {
		email.AddAttachment(s.Attachment)
	}
	s.email = email
	return s
}

func (s *SMailer) Send() error {
	return s.email.Send(s.mailClient)
}
