package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/phaikawl/wade/rt"
	"github.com/phaikawl/wade/test/fntest"
	hm "github.com/phaikawl/wade/test/httpmock"

	c "github.com/phaikawl/wadereddi/common"
)

func TestVoteBox(t *testing.T) {
	score := c.NewScore(69)
	server := hm.NewMock(map[string]hm.Responder{
		"/v": hm.FuncResponder(func(c *hm.Context) hm.Response {
			var vote, lastVote int
			// get values from the query parameters
			fmt.Sscan(c.Request.URL.Query().Get("vote"), &vote)
			fmt.Sscan(c.Request.URL.Query().Get("lastvote"), &lastVote)

			score.UserVote(vote, lastVote)

			return hm.NewOKResponse(fmt.Sprint(score.Score))
		}),
	})

	fntest.NewDummyTestApp("/", server)

	votebox := &VoteBoxModel{}
	votebox.VoteUrl = "/v"
	votebox.Vote = score

	// Start testing
	server.Wait(func() { votebox.DoVote(1) }, 1)
	require.Equal(t, votebox.Vote.Score, 70)
	server.Wait(func() { votebox.DoVote(-1) }, 1)
	require.Equal(t, votebox.Vote.Score, 68)
	server.Wait(func() { votebox.DoVote(-1) }, 1)
	require.Equal(t, votebox.Vote.Score, 69)
}

func startApp(t *testing.T) (myApp *fntest.TestApp) {
	server := hm.NewMock(map[string]hm.Responder{
		"/api/posts":        hm.NewJsonResponse(TestDb),
		"/api/vote/post/3":  hm.NewListResponder([]hm.Responder{hm.NewJsonResponse(97), hm.NewJsonResponse(96)}),
		"/public/*filepath": hm.NewFileResponder("filepath", "../public"),
	})

	myApp = fntest.NewTestApp(rt.Config{}, "/", "../public/index.html", server)

	err := myApp.Start(App{myApp.Application})
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestPostsPage(t *testing.T) {
	app := startApp(t)

	app.GoTo("/posts/top")

	require.Contains(t, app.View.Title(), "Posts")

	posts := app.View.Find("div.post-wrapper")
	require.Contains(t, posts.Eq(0).Text(), "title1")
	require.Contains(t, posts.Eq(1).Text(), "title2")

	voteBtn := app.View.Find("votebox .upvote-btn").Eq(0)
	score := app.View.Find("votebox .score").Eq(0)

	app.View.TriggerEvent(voteBtn, fntest.NewEvent("click"))

	<-app.Render()
	//println("^^", app.Document().Find(".score").Length(), "^^")
	require.Equal(t, score.Text(), "97")

	app.View.TriggerEvent(voteBtn, fntest.NewEvent("click"))
	<-app.Render()
	require.Equal(t, score.Text(), "96")
}

var (
	TestDb = []*c.Post{
		&c.Post{
			Id:     3,
			Author: "poster1",
			Title:  "title1",
			Labels: []string{"label1", "label2"},
			Time:   30,
			Link:   "http://dummy-link.com",
			Comments: []*c.Comment{
				&c.Comment{
					Id:      1,
					Author:  "commenter1",
					Content: "comment1",
					Time:    6,
					Vote:    c.NewScore(3),
				},
				&c.Comment{
					Id:      2,
					Author:  "commenter2",
					Content: "comment2",
					Time:    7,
					Vote:    c.NewScore(4),
				},
			},
			Vote: c.NewScore(96),
		},

		&c.Post{
			Id:      3,
			Author:  "poster2",
			Title:   "title2",
			Labels:  []string{},
			Time:    30,
			Content: "content2",
			Comments: []*c.Comment{
				&c.Comment{
					Id:      3,
					Author:  "commenter1",
					Content: "comment3",
					Time:    6,
					Vote:    c.NewScore(3),
				},
			},
			Vote: c.NewScore(33),
		},
	}
)
