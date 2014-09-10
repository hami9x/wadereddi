package client

import (
	"math/rand"
	"sort"
	"time"

	"github.com/phaikawl/wade"
	"github.com/phaikawl/wade/custom"
	"github.com/phaikawl/wade/taglibs/menu"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const (
	RankModeTop    = "top"
	RankModeLatest = "latest"
)

type (
	Comment struct {
		Author  string
		Content string
		Time    int
		Vote    *Score
	}

	Post struct {
		Id       int
		Title    string
		Author   string
		Labels   []string
		Time     int
		Content  string
		Link     string
		Comments []*Comment
		Vote     *Score
	}

	Score struct {
		Score int
		Voted int
	}

	VoteBoxModel struct {
		custom.BaseProto
		Vote     *Score
		ReSortFn func()
	}

	HomeScope struct {
		*wade.BaseScope
		Posts    []*Post
		SortMode string
	}
)

func (p *Post) Score() int {
	return p.Vote.Score
}

func NewScore(score int) *Score {
	return &Score{score, 0}
}

func (s *Score) VoteUp() {
	s.Score += 1 - s.Voted
	s.Voted = 1
}

func (s *Score) VoteDown() {
	s.Score += -1 - s.Voted
	s.Voted = -1
}

func (m *VoteBoxModel) WrapReSort(voteFn func()) func() {
	return func() {
		voteFn()
		m.ReSortFn()
	}
}

var (
	RandomChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	Database    = []*Post{
		&Post{
			Id:       3,
			Author:   "TheRealDrLee",
			Title:    "Crazy armlet on Techies by Dr.Lee",
			Labels:   []string{"Video", "Top play"},
			Time:     30,
			Link:     "https://www.youtube.com/watch?v=tBYcgpqskdI",
			Comments: GenComments(30),
			Vote:     NewScore(3160),
		},

		&Post{
			Id:       2,
			Author:   "Coward",
			Title:    "Guess I'm going to play cm the next few weeks to dodge Techies",
			Labels:   []string{"Fluff"},
			Time:     10,
			Comments: GenComments(22),
			Vote:     NewScore(100),
			Content:  "I think that's a good idea.",
		},

		&Post{
			Id:       1,
			Author:   "Anticoward",
			Title:    "How to play Techies in ALL modes",
			Labels:   []string{},
			Time:     5,
			Content:  "1. Hack into the servers <br> 2. Inject the code below. <br> 3. Profit!",
			Comments: GenComments(3),
			Vote:     NewScore(100),
		},

		&Post{
			Id:       5,
			Author:   "NigmaNoname",
			Title:    "World's fastest jungle enigma and gem Lv7",
			Labels:   []string{"Video"},
			Time:     40,
			Link:     "https://www.youtube.com/watch?v=iQlFRmVouIA",
			Comments: GenComments(10),
			Vote:     NewScore(99),
		},
	}
)

func NewStrLen(l int) (s string) {
	for i := 0; i < l; i++ {
		s += string(RandomChars[rand.Intn(len(RandomChars))])
	}
	return
}

func GenComments(n int) (ret []*Comment) {
	ret = make([]*Comment, n)
	for i, _ := range ret {
		ret[i] = &Comment{
			Author:  NewStrLen(5),
			Content: NewStrLen(30),
			Time:    rand.Intn(100),
			Vote:    NewScore(rand.Intn(2)),
		}
	}

	return
}

func (home *HomeScope) SortByTop() {
	home.SortMode = RankModeTop
	sort.Sort(PostsByTop(home.Posts))
	home.ApplyChanges(&home.Posts)
}

func (home *HomeScope) SortByLatest() {
	home.SortMode = RankModeLatest
	sort.Sort(PostsByLatest(home.Posts))
	home.ApplyChanges(&home.Posts)
}

func (home *HomeScope) TopRefresh() {
	if home.SortMode == RankModeTop {
		home.SortByTop()
	}
}

func InitFunc(r wade.Registration) {
	r.RegisterDisplayScopes([]wade.PageDesc{
		wade.MakePage("pg-home", "/home", "Home"),
		wade.MakePage("pg-comments", "/comments/:postid", "Comments for %v"),
	}, nil)

	r.RegisterCustomTags(menu.HtmlTags()...)
	r.RegisterCustomTags(custom.HtmlTag{
		Name:       "votebox",
		Attributes: []string{},
		Prototype:  &VoteBoxModel{},
	})

	r.RegisterController("pg-home", func(p *wade.BaseScope) wade.ScopeModel {
		posts := Database

		m := &HomeScope{
			BaseScope: p,
			Posts:     posts,
			SortMode:  RankModeTop,
		}

		m.SortByTop()

		return m
	})
}

// Sorting boilerplates, in real world apps when a database is used, these things
// wouldn't be necessary
type (
	PostsByTop    []*Post
	PostsByLatest []*Post
)

func (a PostsByTop) Len() int           { return len(a) }
func (a PostsByTop) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PostsByTop) Less(i, j int) bool { return a[i].Score() > a[j].Score() }

func (a PostsByLatest) Len() int           { return len(a) }
func (a PostsByLatest) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PostsByLatest) Less(i, j int) bool { return a[i].Time < a[j].Time }
