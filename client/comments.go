package client

import (
	"fmt"

	"github.com/phaikawl/wade/libs/http"
	"github.com/phaikawl/wade/page"
	"github.com/phaikawl/wade/utils"

	c "github.com/phaikawl/wadereddi/common"
)

type (
	// CommentsView is view model for the page pg-comments
	CommentsVM struct {
		httpClient *http.Client
		Post       *c.Post
		RankMode   string
		Comments   []*c.Comment
		NewComment string
	}
)

// AddComment submits the written comment
func (m *CommentsVM) AddComment() {
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

func (vm *CommentsVM) Request(rankMode string) {
	if vm.RankMode == rankMode {
		return
	}

	vm.RankMode = rankMode

	route := fmt.Sprintf("/api/comments/%v", vm.Post.Id)
	url := utils.UrlQuery(route, utils.M{"sort": vm.RankMode})
	r, _ := vm.httpClient.GET(url)
	err := r.ParseJSON(&vm.Comments)
	if err != nil {
		panic(err)
	}
}

func (vm *CommentsVM) postVoteUrl() string {
	return fmt.Sprintf("/api/vote/post/%v", vm.Post.Id)
}

func (vm *CommentsVM) commentVoteUrl(comment *c.Comment) string {
	return fmt.Sprintf("/api/vote/comment/%v", comment.Id)
}

func (am *AppMain) CommentsHandler(ctx page.Context) {
	var postId int
	err := ctx.NamedParams.ScanTo(&postId, "postid")
	if err != nil {
		ctx.GoToPage(PageNotFound)
		return
	}

	// get the post
	var post *c.Post
	r, _ := am.Http.GET(fmt.Sprintf("/api/post/%v", postId))
	err = r.ParseJSON(&post)
	if err != nil {
		panic(err)
	}

	ctx.FormatTitle(post.Title)

	_cvm := &CommentsVM{
		httpClient: am.Http,
		Post:       post,
	}

	_cvm.Request(c.RankModeTop)
}
