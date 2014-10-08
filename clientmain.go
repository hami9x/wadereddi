package main

import (
	"fmt"

	"github.com/phaikawl/wade"
	"github.com/phaikawl/wade/rbackend/clientside"
	"github.com/phaikawl/wadereddi/client"
)

func main() {
	app, err := wade.NewApp(wade.AppConfig{
		BasePath: "/web",
	}, client.InitFunc, clientside.RenderBackend())

	app.Start()

	if err != nil {
		panic(fmt.Sprintf("Failed to load, error: %v.", err.Error()))
	}
}
