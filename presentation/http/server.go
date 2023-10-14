package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/youtrolledhahaha/youdmmmbaa/internal/environment"
	"github.com/youtrolledhahaha/youdmmmbaa/internal/utils/template"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Static("/static", "web/static")
	router.HTMLRender = template.LoadTemplates("web")
	return router
}

func NewServer(router *gin.Engine, configuration *environment.Configuration) error {
	return router.Run(fmt.Sprintf(":%s", configuration.Server.Port))
}
