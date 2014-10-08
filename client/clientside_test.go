package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/phaikawl/wade"
	"github.com/phaikawl/wade/test"
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

	app, err := test.NewDummyApp(t, server)
	if err != nil {
		t.Fatal(err)
	}
	votebox := &VoteBoxModel{}
	app.CustomElemInit(votebox)
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

func TestFunctional(t *testing.T) {
	server := hm.NewMock(map[string]hm.Responder{
		"/api/posts":        hm.NewJsonResponse(TestDb),
		"/public/*filepath": hm.NewFileResponder("filepath", "../public"),
	})

	app, err := test.NewTestApp(t, wade.AppConfig{}, InitFunc, "../public/index.html", server)
	if err != nil {
		panic(err)
	}

	app.Start()
	app.GoTo("/top")
	require.Contains(t, app.View.Title(), "Posts")

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
