package client

import (
	"github.com/phaikawl/wade/page"
	"github.com/phaikawl/wade/rt"
)

type App struct {
	*rt.Application
}

func (am App) Setup(r page.Router) {
	am.autoRouter(r)
}
