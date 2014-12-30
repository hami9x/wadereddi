package client

import (
	"github.com/phaikawl/wade/app"
	"github.com/phaikawl/wade/components/menu"
	"github.com/phaikawl/wade/core"
	"github.com/phaikawl/wade/page"
)

type AppMain struct {
	*app.Application
}

func (am AppMain) Main(app *app.Application) {
	am.Application = app
	r := app.Router()
	r.Handle("/", page.Redirecter{"/posts/top"})
	r.Handle("/posts/:mode", page.Page{
		Id:         PagePosts,
		Title:      "Posts",
		Controller: am.PostsHandler,
	})
	r.Handle("/comments/:postid", page.Page{
		Id:         PageComments,
		Title:      "Comments for %v",
		Controller: am.CommentsHandler,
	})
	r.Otherwise(page.Page{
		Id:    PageNotFound,
		Title: "Page Not Found",
	})

	// Import Wade.Go's standard "w-switchmenu" component
	app.AddComponent(menu.Components()...)

	// Register the "votebox" component
	app.AddComponent(core.ComponentView{
		Name:      "VoteBox",
		Prototype: &VoteBoxModel{},
		Template:  core.HTMLTemplate{"tmpl-votebox"},
	})
}
