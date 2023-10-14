package main

import (
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/client/app"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/client/app/environment"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxx/client/app/ui"
)

var (
	Version       = "dev"
	Port          = ""
	ServerAddress = ""
	Token         = ""
)

func main() {
	ui.ShowMenu(Version, ServerAddress, Port)

	app.New(environment.Load(ServerAddress, Port, Token)).Run()
}
