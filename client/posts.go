package client

import (
	"fmt"

	"github.com/phaikawl/wade/page"
	"github.com/phaikawl/wade/utils"

	c "github.com/phaikawl/wadereddi/common"
)

type PostsVM struct {
	httpClient
	Posts    []*c.Post
	RankMode string
}

func (vm *PostsVM) Request(rankMode string) {
	if vm.RankMode == rankMode {
		return
	}

	vm.RankMode = rankMode
	url := utils.UrlQuery("/api/posts", utils.M{"sort": vm.RankMode})
	r, _ := vm.httpClient.GET(url)
	err := r.ParseJSON(&vm.Posts)
	if err != nil {
		panic(err)
	}
}

func (am AppMain) PostsHandler(ctx page.Context) page.Scope {
	var mode string

	// Get value of the named parameter ":mode" from the url
	ctx.NamedParams.ScanTo(&mode, "mode")

	switch mode {
	case c.RankModeLatest:
		mode = c.RankModeTop
	default:
		mode = c.RankModeLatest
	}

	posts := &PostsVM{
		httpClient: am.Http,
	}

	posts.Request(mode)

	return page.Scope{
		"Pm":        posts,
		"RankModes": c.RankModes,
		"Ctx":       Context{ctx},
		"VoteUrl": func(post *c.Post) string {
			return fmt.Sprintf("/api/vote/post/%v", post.Id)
		},
	}
}
