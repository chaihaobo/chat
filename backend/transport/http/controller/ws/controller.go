package ws

import (
	"net/http"

	"github.com/chaihaobo/gocommon/restkit"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/chaihaobo/chat/application"
	"github.com/chaihaobo/chat/constant"
	"github.com/chaihaobo/chat/resource"
)

type (
	Controller interface {
		Accept(ctx *gin.Context)
	}
	controller struct {
		res resource.Resource
		app application.Application
	}
)

func (c controller) Accept(ctx *gin.Context) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	wsConnection, err := upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	if err != nil {
		c.res.Logger().Error(ctx.Request.Context(), "failed to upgrade websocket", err)
		restkit.HTTPWriteErr(ctx.Writer, constant.ErrSystemMalfunction)
		return
	}
	c.app.Ws().AcceptWsConnection(ctx.Request.Context(), wsConnection)

}

func NewController(res resource.Resource, app application.Application) Controller {
	return &controller{
		res: res,
		app: app,
	}
}
