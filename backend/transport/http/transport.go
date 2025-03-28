package http

import (
	"context"
	"net/http"

	"github.com/chaihaobo/gocommon/constant"
	ginmiddewate "github.com/chaihaobo/gocommon/middleware/http/gin"
	"github.com/chaihaobo/gocommon/restkit"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.uber.org/zap"

	"github.com/chaihaobo/chat/application"
	svcConstant "github.com/chaihaobo/chat/constant"
	"github.com/chaihaobo/chat/model/dto/message"
	"github.com/chaihaobo/chat/model/dto/user"
	"github.com/chaihaobo/chat/resource"
	"github.com/chaihaobo/chat/transport/http/controller"
	"github.com/chaihaobo/chat/transport/http/middleware"
)

type (
	Transport interface {
		Serve() error
		Shutdown() error
	}

	transport struct {
		resource   resource.Resource
		engine     *gin.Engine
		controller controller.Controller
		server     *http.Server
	}
)

func (t *transport) Serve() error {
	var (
		addr = t.resource.Configuration().Service.HTTPPort
		name = t.resource.Configuration().Service.Name
	)
	t.resource.Logger().Info(context.TODO(), "http server started",
		zap.String("name", name),
		zap.String("addr", addr))
	return t.server.ListenAndServe()
}

func (t *transport) Shutdown() error {
	return t.server.Shutdown(context.TODO())
}

func (t *transport) applyRoutes() {
	router := t.engine
	healthController := t.controller.Health()
	userController := t.controller.User()
	wsController := t.controller.Ws()
	messageController := t.controller.Message()
	router.GET("/health", healthController.Health)
	router.GET("/ws", wsController.Accept)

	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", restkit.AdaptToGinHandler(restkit.HandlerFunc[*user.LoginResponse](userController.Login)))
		userGroup.POST("/login/password", restkit.AdaptToGinHandler(restkit.HandlerFunc[*user.LoginResponse](userController.LoginByPassword)))
		userGroup.GET("/friends", restkit.AdaptToGinHandler(restkit.HandlerFunc[user.Users](userController.GetUserFriends)))
		userGroup.GET("/info", restkit.AdaptToGinHandler(restkit.HandlerFunc[*user.User](userController.GetUserInfo)))
	}
	messagesGroup := router.Group("/messages")
	{
		messagesGroup.GET("/recently", restkit.AdaptToGinHandler(restkit.HandlerFunc[*message.GetRecentlyMessagesResponse](messageController.GetFriendRecentlyMessages)))
	}

}

func NewTransport(res resource.Resource, app application.Application) Transport {
	svcConfig := res.Configuration().Service
	gin.SetMode(lo.If(svcConfig.Debug, gin.DebugMode).
		Else(gin.ReleaseMode))
	engine := gin.New()
	constant.MergeServiceErrorCode2HTTPStatus(svcConstant.ServiceErrorCode2HTTPStatus)
	engine.Use(
		ginmiddewate.TelemetryMiddleware(svcConfig.Name, svcConfig.Env, res.Logger()),
		gin.Recovery(),
		middleware.AuthMiddleware(res, app),
	)

	engine.ContextWithFallback = true
	tsp := &transport{
		resource:   res,
		engine:     engine,
		controller: controller.New(res, app),
		server:     &http.Server{Addr: svcConfig.HTTPPort, Handler: engine},
	}
	tsp.applyRoutes()
	return tsp
}
