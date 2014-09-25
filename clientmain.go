package main

import (
	"fmt"

	"github.com/phaikawl/wade"
	"github.com/phaikawl/wade/rbackend/clientside"
	"github.com/phaikawl/wadereddi/client"
)

func main() {
	err := wade.StartApp(wade.AppConfig{
		BasePath: "/web",
	}, client.InitFunc, clientside.RenderBackend())

	if err != nil {
		panic(fmt.Sprintf("Failed to load, error: %v.", err.Error()))
	}
}
