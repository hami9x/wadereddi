package common

import "sort"

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
	Votable interface {
		Voting() *Score
	}

	Comment struct {
		Id      int
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
)

func (lr *ListRank) SortBy(mode string) {
	if lr.RankMode == mode {
		return
	}

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

func (p *Post) Score() int {
	return p.Vote.Score
}

func (c *Comment) Score() int {
	return c.Vote.Score
}

func (p *Post) Voting() *Score {
	return p.Vote
}

func (c *Comment) Voting() *Score {
	return c.Vote
}

func NewScore(score int) *Score {
	return &Score{score, 0}
}

func (s *Score) UserVote(vote int, lastVote int) {
	diff := vote - lastVote
	if diff > 2 {
		diff = -1
	}
	if diff < -2 {
		diff = -2
	}

	s.Score += diff
}

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
