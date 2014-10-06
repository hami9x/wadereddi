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
		custom.BaseProto
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
		// performs an http request to the server to vote, and get the updated
		// score after that
		err := wade.Http().GetJson(&m.Vote.Score, url)
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

	url, err := scope.Url("pg-comments", post.Id)
	if err != nil {
		panic(err)
	}

	return url
}

// AddComment submits the written comment
func (s *CommentsView) AddComment() {
	comment := &c.Comment{
		Author:  CurrentUser,
		Content: s.NewComment,
		Time:    0,
		Vote:    c.NewScore(1),
	}

	go func() {
		// Http request
		wade.Http().POST(
			fmt.Sprintf("/api/comment/add/%v", s.Post.Id),
			comment)

		// Add the comment to the displayed comment list afterwards
		s.Comments = append([]*c.Comment{comment}, s.Comments...)
		s.NewComment = ""
	}()
}

func requestItems(s *wade.Scope, ourl string, rankMode string, listPtr interface{}) (err error) {
	url := wade.UrlQuery(ourl, map[string][]string{
		"sort": []string{rankMode},
	})

	err = wade.Http().GetJson(listPtr, url)

	return
}

func requestPosts(s *wade.Scope, rankMode string, listPtr *[]*c.Post) (err error) {
	return requestItems(s, "/api/posts", rankMode, listPtr)
}

func requestComments(s *wade.Scope, postId int, rankMode string, listPtr *[]*c.Comment) (err error) {
	return requestItems(s, fmt.Sprintf("/api/comments/%v", postId), rankMode, listPtr)
}

func InitFunc(r wade.Registration) {
	// Register the pages
	r.RegisterDisplayScopes([]wade.PageDesc{
		wade.MakePage("pg-posts", "/:mode", "Posts"),
		wade.MakePage("pg-comments", "/comments/:postid", "Comments for %v"),
		wade.MakePage("pg-404", "/notfound", "404 Page Not Found"),
	}, []wade.PageGroupDesc{
		wade.MakePageGroup("grp-main", []string{"pg-posts", "pg-comments"}),
	})

	r.RegisterNotFoundPage("pg-404")

	// Import Wade.Go standard lib's "switchmenu" custom HTML tag
	r.RegisterCustomTags(menu.HtmlTags()...)

	// Register the "votebox" custom HTML tag
	r.RegisterCustomTags(custom.HtmlTag{
		Name:       "votebox",
		Attributes: []string{"Vote", "VoteUrl", "AfterVote"},
		Prototype:  &VoteBoxModel{},
	})

	// Register the page controller for page pg-posts
	r.RegisterController("pg-posts", func(p *wade.Scope) (err error) {
		var mode string
		// Get the named parameter ":mode" from the url
		_ = p.GetParam("mode", &mode)

		switch mode {
		case c.RankModeLatest:
			mode = c.RankModeLatest
		default:
			mode = c.RankModeTop
		}

		// Perform Http request to request the posts
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
		p.AddModel(m)

		// Some minor values and functions used in the HTML code

		p.AddValue("RankModes", c.RankModes)

		p.AddValue("GetLink", func(post *c.Post) string {
			return getLink(p, post)
		})

		p.AddValue("GetVoteUrl", func(post *c.Post) string {
			return fmt.Sprintf("/api/vote/post/%v", post.Id)
		})

		p.AddValue("FetchPosts", func(rankMode string) {
			if m.Rank.RankMode != rankMode {
				m.Rank.RankMode = rankMode
				go requestPosts(p, rankMode, &m.Posts)
			}
		})

		return
	})

	r.RegisterController("pg-comments", func(p *wade.Scope) (err error) {
		var postId int
		err = p.GetParam("postid", &postId)
		if err != nil {
			p.RedirectToPage("pg-404")
			return
		}

		// get the post
		var post *c.Post
		err = wade.Http().GetJson(&post, fmt.Sprintf("/api/post/%v", postId))
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

		p.AddValue("FetchComments", func(rankMode string) {
			if m.RankMode != rankMode {
				m.RankMode = rankMode
				go func() {
					requestComments(p, m.Post.Id, rankMode, &m.Comments)
				}()
			}
		})

		p.AddValue("RankModes", c.RankModes)
		return
	})
}
