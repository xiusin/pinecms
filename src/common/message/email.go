package message

import (
	"bytes"
	"encoding/base64"
	"github.com/kataras/go-mailer"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/config"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strconv"
)

type EmailMessage struct {
	client *mailer.Mailer
}

var ServiceEmailMessage = "pinecms.message.service.email"

func (n *EmailMessage) Init() error {
	conf, err := config.SiteConfig()
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(conf["EMAIL_PORT"])
	if err != nil {
		port = 25
	}
	cfg := mailer.Config{
		Host:      conf["EMAIL_SMTP"],
		Username:  conf["EMAIL_USER"],
		Password:  conf["EMAIL_PWD"],
		Port:      port,
		FromAddr:  conf["EMAIL_EMAIL"],
		FromAlias: conf["EMAIL_SEND_NAME"],
	}

	n.client = mailer.New(cfg)
	return nil
}

func (n *EmailMessage) Notice(receiver []string, params []interface{}, templateId int) error {
	return nil
}

func (n *EmailMessage) Send(receiver []string, subject string, body string) error {
	subject = "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(subject)) + "?="

	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(body)), simplifiedchinese.GBK.NewEncoder()))
	return n.client.SendWithBytes(subject, []byte(data), receiver...)
}

func (n *EmailMessage) UpdateCfg() error { return n.Init() }

func init() {
	di.Set(ServiceEmailMessage, func(builder di.AbstractBuilder) (interface{}, error) {
		email := &EmailMessage{}
		if err := email.Init(); err != nil {
			return nil, err
		}
		return email, nil
	}, true)
}
