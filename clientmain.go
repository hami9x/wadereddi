package main

import (
	"fmt"

	"github.com/phaikawl/wade/app"
	"github.com/phaikawl/wade/rbackend/clientside"
	"github.com/phaikawl/wadereddi/client"
)

func main() {
	app, err := app.New(wade.AppConfig{
		BasePath: "/web",
	}, clientside.CreateBackend())
	
	if err != nil {
		panic(fmt.Sprintf("Failed to load, error: %v.", err.Error()))
	}

	app.Start(client.AppMain{app})
}
