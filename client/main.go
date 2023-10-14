package main

import (
	"github.com/youtrolledhahaha/youdmmmbaa/client/app"
	"github.com/youtrolledhahaha/youdmmmbaa/client/app/environment"
	"github.com/youtrolledhahaha/youdmmmbaa/client/app/ui"
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
