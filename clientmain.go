package main

import (
	"fmt"

	"github.com/phaikawl/wade/platform/clientside"
	"github.com/phaikawl/wade/rt"
	"github.com/phaikawl/wadereddi/client"
)

func main() {
	app := rt.NewApp(rt.Config{
		BasePath: "/web",
	}, clientside.CreateBackend())

	err := app.Start(client.App{app})
	if err != nil {
		panic(fmt.Sprintf("Failed to load, error: %v.", err.Error()))
	}
}
