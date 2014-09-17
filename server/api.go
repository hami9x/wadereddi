package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"path"
	"sort"

	. "github.com/phaikawl/wadereddi/common"
)

func WriteJson(w http.ResponseWriter, data interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	_, err = w.Write(bytes)
	return
}

func handleVote(idstr string, url *neturl.URL, findFn func(id int) (Votable, bool)) (score int, ok bool) {
	var postid int
	fmt.Sscan(idstr, &postid)

	post, ok := findFn(postid)
	if ok {
		var vote, lastVote int
		fmt.Sscan(url.Query().Get("vote"), &vote)
		fmt.Sscan(url.Query().Get("lastvote"), &lastVote)
		post.Voting().UserVote(vote, lastVote)

		score = post.Voting().Score
		return
	}

	return
}

func addHandlers() {
	http.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("sort") {
		case "latest":
			sort.Sort(PostsByLatest(Database))
		default:
			sort.Sort(PostsByTop(Database))
		}

		WriteJson(w, Database)
	})

	postRoute := "/api/post/"
	http.HandleFunc(postRoute, func(w http.ResponseWriter, r *http.Request) {
		if base, postidstr := path.Split(r.URL.Path); base == postRoute {
			var postid int
			fmt.Sscan(postidstr, &postid)

			post, ok := postById(postid)
			if ok {
				WriteJson(w, post)
				return
			}
		}

		http.NotFound(w, r)
	})

	commentsRoute := "/api/comments/"
	http.HandleFunc(commentsRoute, func(w http.ResponseWriter, r *http.Request) {
		if base, postidstr := path.Split(r.URL.Path); base == commentsRoute {
			var postid int
			fmt.Sscan(postidstr, &postid)

			post, ok := postById(postid)
			if ok {
				switch r.URL.Query().Get("sort") {
				case "latest":
					sort.Sort(CommentsByLatest(post.Comments))
				default:
					sort.Sort(CommentsByTop(post.Comments))
				}

				WriteJson(w, post.Comments)
				return
			}
		}

		http.NotFound(w, r)
	})

	pvRoute := "/api/vote/post/"
	http.HandleFunc(pvRoute, func(w http.ResponseWriter, r *http.Request) {
		if base, postidstr := path.Split(r.URL.Path); base == pvRoute {
			score, ok := handleVote(postidstr, r.URL, func(id int) (Votable, bool) {
				return postById(id)
			})

			if ok {
				WriteJson(w, score)
				return
			}
		}

		http.NotFound(w, r)
	})

	cvRoute := "/api/vote/comment/"
	http.HandleFunc(cvRoute, func(w http.ResponseWriter, r *http.Request) {
		if base, idstr := path.Split(r.URL.Path); base == cvRoute {
			score, ok := handleVote(idstr, r.URL, func(id int) (Votable, bool) {
				return commentById(id)
			})

			if ok {
				WriteJson(w, score)
				return
			}
		}

		http.NotFound(w, r)
	})

	caRoute := "/api/comment/add/"
	http.HandleFunc(caRoute, func(w http.ResponseWriter, r *http.Request) {
		if base, postidstr := path.Split(r.URL.Path); base == caRoute {
			var postid int
			fmt.Sscan(postidstr, &postid)
			post, ok := postById(postid)
			if ok {
				bytes, _ := ioutil.ReadAll(r.Body)
				var comment *Comment
				err := json.Unmarshal(bytes, &comment)
				if err != nil {
					http.Error(w, "", http.StatusBadRequest)
					return
				}

				post.Comments = append(post.Comments, comment)
				WriteJson(w, true)
				return
			}
		}

		http.NotFound(w, r)
	})
}
