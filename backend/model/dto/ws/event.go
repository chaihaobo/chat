package ws

// EventSendMessage 发送消息
// EventReceiveMessage  接收消息
const (
	EventSendMessage EventType = iota + 1
	EventReceiveMessage
)

type (
	MessageSend struct {
		To      uint64 `json:"to" mapstructure:"to"`
		Content string `json:"content" mapstructure:"content"`
	}
	MessageReceive struct {
		From    uint64 `json:"from" mapstructure:"from"`
		Content string `json:"content" mapstructure:"content"`
	}
)

type (
	EventType int
)
