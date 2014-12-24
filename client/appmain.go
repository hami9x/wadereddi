package client

import (
	"github.com/phaikawl/wade/app"
	"github.com/phaikawl/wade/components/menu"
	"github.com/phaikawl/wade/core"
	"github.com/phaikawl/wade/page"

	c "github.com/phaikawl/wadereddi/common"
)

const (
	PageComments = "pg-comments"
	PagePosts    = "pg-posts"
	PageNotFound = "pg-404"
	GrpVotable   = "grp-votable"
)

func App() *app.Application {
	return app.App()
}

type AppMain struct {
	*app.Application
}

type Context struct {
	page.Context
}

func (ctx Context) GetLink(post *c.Post) string {
	if post.Link != "" {
		return post.Link
	}

	return ctx.PageUrl(PageComments, post.Id)
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
