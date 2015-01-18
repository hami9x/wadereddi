package client

import (
	"github.com/phaikawl/wade/app"
	"github.com/phaikawl/wade/page"
)

type AppMain struct {
	*app.Application
}

func (am AppMain) Main(app *app.Application) {
	app.PageMgr.SetTemplate(Tmpl_main)
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
}
