package client

import (
	"fmt"

	"github.com/phaikawl/wade/core"
	"github.com/phaikawl/wade/libs/http"
	"github.com/phaikawl/wade/page"
	"github.com/phaikawl/wade/utils"

	c "github.com/phaikawl/wadereddi/common"
)

type (
	// CommentsView is view model for the page pg-comments
	CommentsPageVM struct {
		httpClient *http.Client
		Post       *c.Post
		RankMode   string
		Comments   []*c.Comment
		NewComment string
	}
)

// AddComment submits the written comment
func (m *CommentsPageVM) AddComment() {
	comment := &c.Comment{
		Author:  "me",
		Content: m.NewComment,
		Time:    0,
		Vote:    c.NewScore(1),
	}

	go func() {
		// Http request
		m.httpClient.POST(
			fmt.Sprintf("/api/comment/add/%v", m.Post.Id),
			comment)

		// Add the comment to the displayed comment list afterwards
		m.Comments = append([]*c.Comment{comment}, m.Comments...)
		m.NewComment = ""
	}()
}

func (vm *CommentsPageVM) Request(rankMode string) {
	if vm.RankMode == rankMode {
		return
	}

	vm.RankMode = rankMode

	route := fmt.Sprintf("/api/comments/%v", vm.Post.Id)
	url := utils.UrlQuery(route, utils.Map{"sort": vm.RankMode})
	r, _ := vm.httpClient.GET(url)
	err := r.ParseJSON(&vm.Comments)
	if err != nil {
		panic(err)
	}
}

func (vm *CommentsPageVM) postVoteUrl() string {
	return fmt.Sprintf("/api/vote/post/%v", vm.Post.Id)
}

func (vm *CommentsPageVM) commentVoteUrl(comment *c.Comment) string {
	return fmt.Sprintf("/api/vote/comment/%v", comment.Id)
}

func (am App) CommentsPageHandler(ctx *page.Context) *core.VNode {
	var postId int
	err := ctx.NamedParams.ScanTo(&postId, "postid")
	if err != nil {
		return ctx.GoToPage(NotFoundPage)
	}

	// get the post
	var post *c.Post
	r, _ := am.Http.GET(fmt.Sprintf("/api/post/%v", postId))
	err = r.ParseJSON(&post)
	if err != nil {
		panic(err)
	}

	ctx.FormatTitle(post.Title)

	vm := &CommentsPageVM{
		httpClient: am.Http,
		Post:       post,
	}

	vm.Request(c.RankModeTop)
	return vm.Template()
}
