package client

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/phaikawl/wade"
	"github.com/phaikawl/wade/custom"
	"github.com/phaikawl/wade/taglibs/menu"

	c "github.com/phaikawl/wadereddi/common"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	CurrentUser = "Somebody"
)

type (
	// VoteBoxModel is the prototype for the "votebox" custom element
	VoteBoxModel struct {
		*wade.BaseProto
		Vote      *c.Score
		VoteUrl   string
		AfterVote func() // function to be called after a vote is done
	}

	// PostsView is view model for the page pg-posts
	PostsView struct {
		s     *wade.Scope
		Posts []*c.Post
		Rank  *c.ListRank
	}

	// CommentsView is view model for the page pg-comments
	CommentsView struct {
		s          *wade.Scope
		Post       *c.Post
		RankMode   string
		Comments   []*c.Comment
		NewComment string
	}
)

// DoVote performs the vote with the given point changed (-1 = vote down, 1 = vote up)
func (m *VoteBoxModel) DoVote(vote int) {
	if m.VoteUrl == "" {
		panic("VoteUrl has not been set.")
	}

	if vote < -1 || vote > 1 {
		panic("Illegal vote value.")
	}

	lastVote := m.Vote.Voted
	m.Vote.Voted = vote
	if vote == lastVote {
		vote, lastVote = -lastVote, 0
		m.Vote.Voted = 0
	}

	url := wade.UrlQuery(m.VoteUrl, map[string][]string{
		"vote":     []string{fmt.Sprintf("%v", vote)},
		"lastvote": []string{fmt.Sprintf("%v", lastVote)},
	})

	// http request is blocking, so we put it in a goroutine, typical Go
	go func() {
		// performs an http request to the server to vote, and assign the updated score
		// to m.Vote.Score after that
		err := m.BaseProto.App.Http().GetJson(&m.Vote.Score, url)
		if err != nil {
			return
		}

		if m.AfterVote != nil {
			m.AfterVote()
		}
	}()
}

func getLink(scope *wade.Scope, post *c.Post) string {
	if post.Link != "" {
		return post.Link
	}

	url, err := scope.GetUrl("pg-comments", post.Id)
	if err != nil {
		panic(err)
	}

	return url
}

// AddComment submits the written comment
func (m *CommentsView) AddComment() {
	comment := &c.Comment{
		Author:  CurrentUser,
		Content: m.NewComment,
		Time:    0,
		Vote:    c.NewScore(1),
	}

	go func() {
		// Http request
		m.s.Http().POST(
			fmt.Sprintf("/api/comment/add/%v", m.Post.Id),
			comment)

		// Add the comment to the displayed comment list afterwards
		m.Comments = append([]*c.Comment{comment}, m.Comments...)
		m.NewComment = ""
	}()
}

func requestPosts(s *wade.Scope, rankMode string, listPtr *[]*c.Post) (err error) {
	return requestItems(s, "/api/posts", rankMode, listPtr)
}

func requestComments(s *wade.Scope, postId int, rankMode string, listPtr *[]*c.Comment) (err error) {
	return requestItems(s, fmt.Sprintf("/api/comments/%v", postId), rankMode, listPtr)
}

func (m *CommentsView) FetchComments(rankMode string) {
	if m.RankMode != rankMode {
		m.RankMode = rankMode
		go requestComments(m.s, m.Post.Id, rankMode, &m.Comments)
	}
}

func (m *PostsView) FetchPosts(rankMode string) {
	if m.Rank.RankMode != rankMode {
		m.Rank.RankMode = rankMode
		go requestPosts(m.s, rankMode, &m.Posts)
	}
}

func requestItems(s *wade.Scope, ourl string, rankMode string, listPtr interface{}) (err error) {
	url := wade.UrlQuery(ourl, map[string][]string{
		"sort": []string{rankMode},
	})

	err = s.Http().GetJson(listPtr, url)

	return
}

func AppFunc(app *wade.Application) {
	app.SetStartPath("/posts/top")

	// Declare the pages
	app.Register.DisplayScopes([]wade.PageDesc{
		wade.MakePage("pg-posts", "/posts/:mode", "Posts"),
		wade.MakePage("pg-comments", "/comments/:postid", "Comments for %v"),
		wade.MakePage("pg-404", "", "404 Page Not Found"),
	}, []wade.PageGroupDesc{
		wade.MakePageGroup("grp-main", []string{"pg-posts", "pg-comments"}),
	})

	app.Register.NotFoundPage("pg-404")

	// Import Wade.Go's standard "switchmenu" custom HTML tag
	app.Register.CustomTags(menu.HtmlTags()...)

	// Register the "votebox" custom HTML tag
	app.Register.CustomTags(custom.HtmlTag{
		Name:       "votebox",
		Attributes: []string{"Vote", "VoteUrl", "AfterVote"},
		Prototype:  &VoteBoxModel{},
	})

	// Register the page controller for page pg-posts
	app.Register.Controller("pg-posts", func(p *wade.Scope) (err error) {
		var mode string
		// Get value of the named parameter ":mode" from the url
		_ = p.NamedParams.GetTo("mode", &mode)

		switch mode {
		case c.RankModeLatest:
			mode = c.RankModeLatest
		default:
			mode = c.RankModeTop
		}

		// Perform Http request to request the posts
		// Notice:
		// We don't use and shouldn't use a separate goroutine for Http request in a
		// page controller. Just call it directly,
		// because each controller is run in its own goroutine already by Wade, and
		// the page is rendered right after this function returns. If some content
		// is not available when this function returns, it will not get displayed.
		var posts []*c.Post
		err = requestPosts(p, mode, &posts)
		if err != nil {
			return
		}

		m := &PostsView{
			s:     p,
			Posts: posts,
			Rank: &c.ListRank{
				RankMode: mode,
				List:     c.PostsRank{posts},
			},
		}

		// Add the view model
		// all fields of the model are then available in the HTML code for binding
		p.AddModel(m)

		// Below are some minor values and helper functions used in the HTML code
		// These things don't have anything to do with the logic, flow
		// or changing the data, so using them this way is more
		// convenient without any real downsides

		p.AddValue("RankModes", c.RankModes)

		p.AddValue("GetLink", func(post *c.Post) string {
			return getLink(p, post)
		})

		p.AddValue("GetVoteUrl", func(post *c.Post) string {
			return fmt.Sprintf("/api/vote/post/%v", post.Id)
		})

		return
	})

	app.Register.Controller("pg-comments", func(p *wade.Scope) (err error) {
		var postId int
		err = p.NamedParams.GetTo("postid", &postId)
		if err != nil {
			p.GoToPage("pg-404")
			return
		}

		// get the post
		var post *c.Post
		err = p.Http().GetJson(&post, fmt.Sprintf("/api/post/%v", postId))
		if err != nil {
			return
		}

		var comments []*c.Comment
		err = requestComments(p, postId, c.RankModeTop, &comments)
		if err != nil {
			return
		}

		p.FormatTitle(post.Title)

		m := &CommentsView{
			s:        p,
			Post:     post,
			Comments: comments,
			RankMode: c.RankModeTop,
		}

		p.AddModel(m)

		p.AddValue("GetLink", func(post *c.Post) string {
			return getLink(p, post)
		})

		p.AddValue("GetCommentVoteUrl", func(comment *c.Comment) string {
			return fmt.Sprintf("/api/vote/comment/%v", comment.Id)
		})

		p.AddValue("GetPostVoteUrl", func(post *c.Post) string {
			return fmt.Sprintf("/api/vote/post/%v", post.Id)
		})

		p.AddValue("RankModes", c.RankModes)
		return
	})
}
