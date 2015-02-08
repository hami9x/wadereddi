package client

import (
	"fmt"

	"github.com/phaikawl/wade/libs/http"
	"github.com/phaikawl/wade/page"
	"github.com/phaikawl/wade/utils"
	"github.com/phaikawl/wade/vdom"

	c "github.com/phaikawl/wadereddi/common"
)

type PostsPageVM struct {
	httpClient *http.Client
	Posts      []*c.Post
	RankMode   string
}

var posts []*c.Post

func (vm *PostsPageVM) Request(rankMode string) {
	if vm.RankMode == rankMode {
		return
	}

	vm.RankMode = rankMode
	url := utils.UrlQuery("/api/posts", utils.Map{"sort": vm.RankMode})
	r, _ := vm.httpClient.GET(url)
	err := r.ParseJSON(&vm.Posts)

	if err != nil {
		panic(err)
	}
}

func (vm *PostsPageVM) voteUrl(post *c.Post) string {
	return fmt.Sprintf("/api/vote/post/%v", post.Id)
}

func (am App) PostsPageHandler(ctx *page.Context) *vdom.VNode {
	var mode string

	// Get value of the named parameter ":mode" from the url
	ctx.NamedParams.ScanTo(&mode, "mode")

	switch mode {
	case c.RankModeLatest:
		mode = c.RankModeLatest
	default:
		mode = c.RankModeTop
	}

	// Export to view symbol
	vm := &PostsPageVM{
		httpClient: am.Http,
	}

	vm.Request(mode)
	posts = vm.Posts
	return vm.Template()
}
