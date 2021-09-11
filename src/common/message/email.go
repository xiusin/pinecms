package message

import (
	"github.com/kataras/go-mailer"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/config"
	"strconv"
)

type EmailMessage struct{
	client *mailer.Mailer
}

var ServiceEmailMessage = "pinecms.message.service.email"

func (n *EmailMessage) Init() error {
	conf,err := config.SiteConfig()
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(conf["EMAIL_PORT"])
	if err != nil {
		port = 25
	}
	n.client = mailer.New(mailer.Config{
		Host:      conf["EMAIL_SMTP"],
		Username:  conf["EMAIL_USER"],
		Password:  conf["EMAIL_PWD"],
		Port:      port,
		FromAlias: conf["EMAIL_SEND_NAME"],
	})
	return nil
 }

func (n *EmailMessage) Notice(receiver []string, params []interface{}, templateId int) error { return nil }

func (n *EmailMessage) Send(receiver []string, msg string, typo int) error {
	return n.client.Send("","", receiver...)
}

func (n *EmailMessage) UpdateCfg() error {
	return n.Init()
}

func init() {
	di.Set(ServiceEmailMessage, func(builder di.AbstractBuilder) (interface{}, error) {
		email := &EmailMessage{}
		if err := email.Init(); err != nil {
			return nil, err
		}
		return email, nil
	}, true)
}
