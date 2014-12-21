package common

type RankModeInfo struct {
	Code string
	Name string
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
)

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
		diff = 2
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
