package client

import (
	"fmt"

	"github.com/phaikawl/wade/page"
	"github.com/phaikawl/wade/utils"

	c "github.com/phaikawl/wadereddi/common"
)

type (
	// CommentsView is view model for the page pg-comments
	CommentsVM struct {
		httpClient
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
		App().Http.POST(
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

func (am *AppMain) CommentsHandler(ctx page.Context) page.Scope {
	var postId int
	err := ctx.NamedParams.ScanTo(&postId, "postid")
	if err != nil {
		ctx.GoToPage(PageNotFound)
		return nil
	}

	// get the post
	var post *c.Post
	r, _ := am.Http.GET(fmt.Sprintf("/api/post/%v", postId))
	err = r.ParseJSON(&post)
	if err != nil {
		panic(err)
	}

	comments := &CommentsVM{
		httpClient: am.Http,
		Post:       post,
	}

	comments.Request(c.RankModeTop)

	ctx.FormatTitle(post.Title)

	return page.Scope{
		"Cm":        comments,
		"RankModes": c.RankModes,
		"Ctx":       Context{ctx},
		"CommentVoteUrl": func(comment *c.Comment) string {
			return fmt.Sprintf("/api/vote/comment/%v", comment.Id)
		},
		"PostVoteUrl": fmt.Sprintf("/api/vote/post/%v", post.Id),
	}
}
