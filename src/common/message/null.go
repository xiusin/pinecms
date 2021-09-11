package message

import "github.com/xiusin/pine/di"

type NullMessage struct{}

var ServiceNullMessage = "pinecms.message.service.null"

func (n NullMessage) Init() error { return nil }

func (n NullMessage) Notice(receiver []string, params []interface{}, templateId int) error { return nil }

func (n NullMessage) Send(receiver []string, msg string, typo int) error { return nil }

func (n NullMessage) UpdateCfg() error { return nil }

func init() {
	di.Set(ServiceSmsMessage, func(builder di.AbstractBuilder) (interface{}, error) {
		return &NullMessage{}, nil
	}, true)
}
