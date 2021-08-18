package message

// AbstractMessage 发送接口
type AbstractMessage interface {
	Send(cfg map[string]interface{}) error
}

type NullMessage struct{}

func (n NullMessage) Send(_ map[string]interface{}) error {
	return nil
}
