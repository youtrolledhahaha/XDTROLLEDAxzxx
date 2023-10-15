package main

import (
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/environment"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxclient/app/ui"
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
