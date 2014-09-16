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

type RankModeInfo struct {
	Name          string
	DisplayedName string
}

const (
	RankModeTop    = "top"
	RankModeLatest = "latest"
)

var (
	RankModes = []RankModeInfo{
		RankModeInfo{RankModeTop, "Top"},
		RankModeInfo{RankModeLatest, "Latest"},
	}
)

var (
	CurrentUser = "Somebody"

	PostsSort = map[string]func(posts []*Post){
		RankModeTop: func(posts []*Post) {
			sort.Sort(PostsByTop(posts))
		},
		RankModeLatest: func(posts []*Post) {
			sort.Sort(PostsByLatest(posts))
		},
	}

	CommentsSort = map[string]func(comments []*Comment){
		RankModeTop: func(comments []*Comment) {
			sort.Sort(CommentsByTop(comments))
		},
		RankModeLatest: func(comments []*Comment) {
			sort.Sort(CommentsByLatest(comments))
		},
	}
)

type (
	RankedList interface {
		Sort(rankMode string)
	}

	ListRank struct {
		RankMode string
		List     RankedList
	}

	PostsRank struct {
		List []*Post
	}

	CommentsRank struct {
		List []*Comment
	}

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
		Rank      *ListRank
		Posts     []*Post
		RankModes []RankModeInfo
	}

	CommentsScope struct {
		*wade.BaseScope
		Rank      *ListRank
		Post      *Post
		RankModes []RankModeInfo
	}
)

func (p *Post) Score() int {
	return p.Vote.Score
}

func (c *Comment) Score() int {
	return c.Vote.Score
}

func NewScore(score int) *Score {
	return &Score{score, 0}
}

func (s *Score) VoteUp() {
	switch s.Voted {
	case -1:
		s.Score += 2
		s.Voted = 1
	case 0:
		s.Score++
		s.Voted = 1
	case 1:
		s.Score--
		s.Voted = 0
	}
}

func (s *Score) VoteDown() {
	switch s.Voted {
	case 1:
		s.Score -= 2
		s.Voted = -1
	case 0:
		s.Score--
		s.Voted = -1
	case -1:
		s.Score++
		s.Voted = 0
	}
}

func (m *VoteBoxModel) ReSort(voteFn func()) {
	voteFn()
	if m.ReSortFn != nil {
		m.ReSortFn()
	}
}

var (
	RandomChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789            "
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
			Author:  NewStrLen(3 + rand.Intn(5)),
			Content: NewStrLen(rand.Intn(8 + 50)),
			Time:    rand.Intn(100),
			Vote:    NewScore(rand.Intn(2)),
		}
	}

	return
}

func (lr *ListRank) SortBy(mode string) {
	lr.RankMode = mode
	lr.List.Sort(mode)
}

func (lr *ListRank) TopRefresh() {
	if lr.RankMode == RankModeTop {
		lr.List.Sort(RankModeTop)
	}
}

func (r PostsRank) Sort(rankMode string) {
	PostsSort[rankMode](r.List)
}

func (r CommentsRank) Sort(rankMode string) {
	CommentsSort[rankMode](r.List)
}

func postById(id int) (p *Post, ok bool) {
	for _, post := range Database {
		if post.Id == id {
			return post, true
		}
	}

	return
}

func getLink(scope *wade.BaseScope, post *Post) string {
	if post.Link != "" {
		return post.Link
	}

	url, err := scope.Url("pg-comments", post.Id)
	if err != nil {
		panic(err)
	}

	return url
}

func (s *HomeScope) GetLink(post *Post) string {
	return getLink(s.BaseScope, post)
}

func (s *CommentsScope) GetLink(post *Post) string {
	return getLink(s.BaseScope, post)
}

func (s *CommentsScope) AddItem() {
	s.Post.Comments = append(s.Post.Comments, &Comment{Author: "abc", Content: "def", Time: 0, Vote: NewScore(0)})
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
		Attributes: []string{"Vote", "ReSortFn"},
		Prototype:  &VoteBoxModel{},
	})

	r.RegisterController("pg-home", func(p *wade.BaseScope) wade.ScopeModel {
		posts := Database

		m := &HomeScope{
			BaseScope: p,
			Posts:     posts,
			Rank: &ListRank{
				RankMode: RankModeTop,
				List:     PostsRank{posts},
			},
			RankModes: RankModes,
		}

		m.Rank.SortBy(m.Rank.RankMode)

		return m
	})

	r.RegisterController("pg-comments", func(p *wade.BaseScope) wade.ScopeModel {
		var postId int
		err := p.GetParam("postid", &postId)
		if err != nil {
			p.RedirectToPage("pg-404")
		}

		post, ok := postById(postId)
		if !ok {
			p.RedirectToPage("pg-404")
		}

		p.FormatTitle(post.Title)

		m := &CommentsScope{
			BaseScope: p,
			Post:      post,
			Rank: &ListRank{
				RankMode: RankModeTop,
				List:     CommentsRank{post.Comments},
			},
			RankModes: RankModes,
		}

		m.Rank.SortBy(m.Rank.RankMode)

		return m
	})
}

// Sorting boilerplates, in real world apps when a database is used, these things
// wouldn't be necessary
type (
	PostsByTop       []*Post
	PostsByLatest    []*Post
	CommentsByTop    []*Comment
	CommentsByLatest []*Comment
)

func (a PostsByTop) Len() int           { return len(a) }
func (a PostsByTop) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PostsByTop) Less(i, j int) bool { return a[i].Score() > a[j].Score() }

func (a PostsByLatest) Len() int           { return len(a) }
func (a PostsByLatest) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PostsByLatest) Less(i, j int) bool { return a[i].Time < a[j].Time }

func (a CommentsByTop) Len() int           { return len(a) }
func (a CommentsByTop) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CommentsByTop) Less(i, j int) bool { return a[i].Score() > a[j].Score() }

func (a CommentsByLatest) Len() int           { return len(a) }
func (a CommentsByLatest) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CommentsByLatest) Less(i, j int) bool { return a[i].Time < a[j].Time }
