package main

import (
	"fmt"

	"github.com/phaikawl/wade/rt"
	"github.com/phaikawl/wade/platform/clientside"
	"github.com/phaikawl/wadereddi/client"
)

func main() {
	ap := rt.NewApp(app.Config{
		BasePath: "/web",
	}, clientside.CreateBackend())

	err := ap.Start(client.App{ap})
	if err != nil {
		panic(fmt.Sprintf("Failed to load, error: %v.", err.Error()))
	}
}
