package message

import (
	"context"

	"github.com/samber/lo"

	"github.com/chaihaobo/chat/constant"
	"github.com/chaihaobo/chat/infrastructure/store/repository/querytypes"
	"github.com/chaihaobo/chat/model/entity"
	"github.com/chaihaobo/chat/tools"
)

type (
	Message struct {
		ID      uint64 `json:"id"`
		Content string `json:"content"`
		From    uint64 `json:"from"`
		To      uint64 `json:"to"`
	}

	GetRecentlyMessagesResponse struct {
		HasMore  bool       `json:"has_more"`
		Messages []*Message `json:"messages"`
	}

	GetRecentlyMessagesRequest struct {
		FriendUserID uint64 `json:"friend_user_id" form:"friend_user_id" validate:"required"`
		Offset       int    `json:"offset" form:"offset"`
		Limit        int    `json:"limit" form:"limit"`
	}
)

func (g *GetRecentlyMessagesRequest) FillDefault() {
	if g.Offset < 0 {
		g.Offset = 0
	}
	if g.Limit <= 0 {
		g.Limit = constant.PaginationDefaultLimit
	}

	if g.Limit > constant.PaginationMaxLimit {
		g.Limit = constant.PaginationMaxLimit
	}

}

func (g *GetRecentlyMessagesRequest) ToQuery(ctx context.Context) *querytypes.RecentlyMessageQuery {
	return &querytypes.RecentlyMessageQuery{
		Pagination: querytypes.Pagination{
			Offset: g.Offset,
			Limit:  g.Limit,
		},
		FriendUserID: lo.ToPtr(g.FriendUserID),
		UserID:       lo.ToPtr(tools.ContextUserID(ctx)),
	}

}

func NewGetRecentlyMessagesResponse(request *GetRecentlyMessagesRequest, messages entity.Messages, total int64) *GetRecentlyMessagesResponse {
	return &GetRecentlyMessagesResponse{
		HasMore: request.Offset+request.Limit+1 < int(total),
		Messages: lo.Map(messages, func(item *entity.Message, index int) *Message {
			return &Message{
				ID:      item.ID,
				Content: item.Content,
				From:    item.From,
				To:      item.To,
			}
		}),
	}
}
