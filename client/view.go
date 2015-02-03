package client

import (
	"github.com/phaikawl/wade/page"
	"github.com/phaikawl/wade/rt"
	c "github.com/phaikawl/wadereddi/common"
)

var (
	RankModes = c.RankModes
)

type context struct {
	*page.Context
}

func app() *rt.Application {
	return rt.App()
}

func ctx() context {
	return context{app().PageMgr.Context()}
}

func (ctx context) getPostLink(post *c.Post) string {
	if post.Link != "" {
		return post.Link
	}

	return ctx.PageUrl(CommentsPage, post.Id)
}
