package message

import "github.com/chaihaobo/chat/constant"

type (
	Message struct {
		Content string `json:"content"`
		From    uint64 `json:"from"`
		To      uint64 `json:"to"`
	}

	GetRecentlyMessagesResponse struct {
		HasMore  bool       `json:"has_more"`
		Messages []*Message `json:"messages"`
	}

	GetRecentlyMessagesRequest struct {
		FriendUserID uint64 `json:"friend_user_id" validate:"friend_user_id"`
		Offset       int    `json:"offset"`
		Limit        int    `json:"limit"`
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
