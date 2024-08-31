package http

import (
	"dddapib/internal/domain/service"
	"dddapib/internal/infrastructure/transport/http/handler/task"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve(svc *service.Service) {
	router := gin.Default()
	router.Use(gin.Recovery())

	task.Init(router, svc)
	// add more API here

	endless.ListenAndServe(viper.GetString("infra.http.addr"), router)
}
