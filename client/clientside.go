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
	VoteBoxModel struct {
		custom.BaseProto
		Vote      *c.Score
		VoteUrl   string
		AfterVote func()
	}

	HomeView struct {
		s     *wade.Scope
		Posts []*c.Post
		Rank  *c.ListRank
	}

	CommentsView struct {
		s          *wade.Scope
		Post       *c.Post
		RankMode   string
		Comments   []*c.Comment
		NewComment string
	}
)

func (m *VoteBoxModel) DoVote(vote int) {
	go func() {
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

		resp, err := wade.Http().GET(url)
		if err != nil || resp.Failed() {
			return
		}

		resp.DecodeTo(&m.Vote.Score)

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

func (s *CommentsView) AddComment() {
	go func() {
		comment := &c.Comment{
			Author:  CurrentUser,
			Content: s.NewComment,
			Time:    0,
			Vote:    c.NewScore(1),
		}

		resp, err := wade.Http().POST(
			fmt.Sprintf("/api/comment/add/%v", s.Post.Id),
			comment)

		if err != nil || resp.Failed() {
			return
		}

		s.Comments = append([]*c.Comment{comment}, s.Comments...)
		s.NewComment = ""
	}()
}

func requestItems(s *wade.Scope, ourl string, rankMode string, listPtr interface{}) (err error) {
	url := wade.UrlQuery(ourl, map[string][]string{
		"sort": []string{rankMode},
	})

	resp, err := wade.Http().GET(url)
	if err != nil || resp.Failed() {
		return
	}

	resp.DecodeTo(listPtr)
	return
}

func requestPosts(s *wade.Scope, rankMode string, listPtr *[]*c.Post) (err error) {
	return requestItems(s, "/api/posts", rankMode, listPtr)
}

func requestComments(s *wade.Scope, postId int, rankMode string, listPtr *[]*c.Comment) (err error) {
	return requestItems(s, fmt.Sprintf("/api/comments/%v", postId), rankMode, listPtr)
}

func InitFunc(r wade.Registration) {
	r.RegisterDisplayScopes([]wade.PageDesc{
		wade.MakePage("pg-home", "/home", "Home"),
		wade.MakePage("pg-comments", "/comments/:postid", "Comments for %v"),
		wade.MakePage("pg-404", "/notfound", "404 Page Not Found"),
	}, []wade.PageGroupDesc{
		wade.MakePageGroup("grp-main", []string{"pg-home", "pg-comments"}),
	})

	r.RegisterNotFoundPage("pg-404")

	r.RegisterCustomTags(menu.HtmlTags()...)

	r.RegisterCustomTags(custom.HtmlTag{
		Name:       "votebox",
		Attributes: []string{"Vote", "VoteUrl", "AfterVote"},
		Prototype:  &VoteBoxModel{},
	})

	r.RegisterController("pg-home", func(p *wade.Scope) (err error) {
		var posts []*c.Post
		err = requestPosts(p, c.RankModeTop, &posts)
		if err != nil {
			return
		}

		m := &HomeView{
			s:     p,
			Posts: posts,
			Rank: &c.ListRank{
				RankMode: c.RankModeTop,
				List:     c.PostsRank{posts},
			},
		}

		p.AddModel(m)

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
				go func() {
					requestPosts(p, rankMode, &m.Posts)
				}()
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
		resp, err := wade.Http().GET(fmt.Sprintf("/api/post/%v", postId))
		if err != nil || resp.Failed() {
			return
		}

		var post *c.Post
		err = resp.DecodeTo(&post)
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
