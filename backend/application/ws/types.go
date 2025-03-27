package ws

import (
	"time"

	userdto "github.com/chaihaobo/chat/model/dto/user"
	"github.com/chaihaobo/chat/model/dto/ws"
	"github.com/chaihaobo/chat/model/entity"
	"github.com/chaihaobo/chat/tools"
)

type (
	// message 是内部传递的消息结构体, 用于和广播到各个节点
	message struct {
		From    *userdto.User `json:"from"`
		To      *userdto.User `json:"to"`
		Content string        `json:"content"`
	}
)

func (m message) toEntity() *entity.Message {
	return &entity.Message{
		BaseEntity: entity.BaseEntity{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		From:    m.From.ID,
		To:      m.To.ID,
		Content: m.Content,
	}
}

func newMessage(con *Connection, request *ws.MessageSend) *message {
	return &message{
		From: &userdto.User{
			ID:       tools.ContextUserID(con.Context()),
			UserName: tools.ContextUserName(con.Context()),
			Avatar:   tools.ContextUserAvatar(con.Context()),
		},
		To: &userdto.User{
			ID: request.To,
		},
		Content: request.Content,
	}
}
