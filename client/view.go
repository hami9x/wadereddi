package client

import (
	"github.com/phaikawl/wade/app"
	"github.com/phaikawl/wade/page"
	c "github.com/phaikawl/wadereddi/common"
)

var (
	_rankModes = c.RankModes
	_app       = app.App()
)

var (
	_cvm *CommentsVM
	_pvm *PostsVM
)

const (
	PageComments = "pg-comments"
	PagePosts    = "pg-posts"
	PageNotFound = "pg-404"
	GrpVotable   = "grp-votable"
)

type context struct {
	page.Context
}

func ctx() context {
	return context{_app.PageMgr.Context()}
}

func (ctx context) getPostLink(post *c.Post) string {
	if post.Link != "" {
		return post.Link
	}

	return ctx.PageUrl(PageComments, post.Id)
}
