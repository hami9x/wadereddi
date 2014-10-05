package main

import (
	"math/rand"
	. "github.com/phaikawl/wadereddi/common"
)

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
			Comments: GenComments(55),
			Vote:     NewScore(3160),
		},

		&Post{
			Id:       2,
			Author:   "Coward",
			Title:    "Guess I'm going to play Captain's Mode the next few weeks to dodge Techies",
			Labels:   []string{"Fluff"},
			Time:     10,
			Comments: GenComments(30),
			Vote:     NewScore(100),
			Content:  "I think that's a good idea.",
		},

		&Post{
			Id:       1,
			Author:   "Anticoward",
			Title:    "How to play Techies in ALL modes",
			Labels:   []string{},
			Time:     5,
			Content:  "I know u guys are crazy about Techies, and I'm here to give you guys... <br> A troll post XD.",
			Comments: GenComments(5),
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
	s += string(RandomChars[rand.Intn(len(RandomChars))])
	for i := 0; i < l; i++ {
		s += string((RandomChars + "              ")[rand.Intn(len(RandomChars)+10)])
	}
	return
}

func GenComments(n int) (ret []*Comment) {
	ret = make([]*Comment, n)
	for i, _ := range ret {
		ret[i] = &Comment{
			Id:      rand.Intn(32767),
			Author:  NewStrLen(3 + rand.Intn(5)),
			Content: NewStrLen(rand.Intn(8 + 50)),
			Time:    rand.Intn(100),
			Vote:    NewScore(rand.Intn(2)),
		}
	}

	return
}

func postById(id int) (p *Post, ok bool) {
	for _, post := range Database {
		if post.Id == id {
			return post, true
		}
	}

	return
}

func commentById(id int) (p *Comment, ok bool) {
	for _, post := range Database {
		for _, comment := range post.Comments {
			if comment.Id == id {
				return comment, true
			}
		}
	}

	return
}
