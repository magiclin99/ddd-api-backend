package http

import (
	"dddapib/internal/domain/service"
	"dddapib/internal/infrastructure/transport/http/handler/task"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server interface {
	ListenAndServe() error
	Close() error
}

func NewServer(svc *service.Service) Server {
	router := gin.Default()
	router.Use(gin.Recovery())

	task.Init(router, svc)
	// add more API here

	return endless.NewServer(viper.GetString("infra.http.addr"), router)
}
