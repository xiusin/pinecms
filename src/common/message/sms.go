package message

//import (
//	"sync"
//
//	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
//	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
//	"github.com/alibabacloud-go/tea/tea"
//	"xorm.io/xorm"
//	"github.com/xiusin/pine/di"
//)
//
const ServiceSmsMessage = "pinecms.message.service.sms"
//
//type SmsMessage struct {
//	client *dysmsapi20170525.Client
//	orm *xorm.Engine
//	sync.Mutex
//}
//
//func (n *SmsMessage) UpdateCfg() error {
//	return n.Init()
//}
//
//func (n *SmsMessage) Init() error {
//	n.Lock()
//	defer n.Unlock()
//	var err error
//	var client *dysmsapi20170525.Client
//	config := &openapi.Config{
//		AccessKeyId: nil,
//		AccessKeySecret: nil,
//	}
//	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
//	if client, err = dysmsapi20170525.NewClient(config); err == nil {
//		n.client = client
//	}
//	_orm , err := di.Get(&xorm.Engine{})
//	if err != nil {
//		return err
//	}
//	n.orm = _orm.(*xorm.Engine)
//	return err
//}
//
//func (n *SmsMessage) Notice(receiver []string, params []interface{}, templateId int) error {
//	return nil
//}
//
//func (n *SmsMessage) Send(receiver []string, msg string, typo int) error {
//
//	for _, v := range receiver {
//		sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
//			PhoneNumbers: tea.String(v),
//		}
//		if _, err := n.client.SendSms(sendSmsRequest); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func init() {
//	di.Set(ServiceSmsMessage, func(builder di.AbstractBuilder) (interface{}, error) {
//		msgService := &SmsMessage{}
//		err := msgService.Init()
//		if err != nil {
//			return nil, err
//		}
//		return msgService, nil
//	}, true)
//}
