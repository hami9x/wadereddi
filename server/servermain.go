package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	wadeserv "github.com/phaikawl/wade/platform/serverside"
	"github.com/phaikawl/wade/rt"
	"github.com/phaikawl/wadereddi/client"
)

const (
	ServersidePrerender = false
	DevMode             = true
)

func main() {
	addHandlers()

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("../public/"))))
	if DevMode {
		http.Handle("/gopath/", http.StripPrefix("/gopath", http.FileServer(http.Dir(os.Getenv("GOPATH")))))
	}

	http.HandleFunc("/web/", func(w http.ResponseWriter, r *http.Request) {
		indexBytes, err := ioutil.ReadFile("../public/index.html")
		if err != nil {
			panic(err)
		}

		if !ServersidePrerender {
			w.Header().Set("Content-Type", "text/html")
			w.Write(indexBytes)
		} else {
			httpBkn := wadeserv.NewHttpBackend(http.DefaultServeMux, r, "/api")
			app := wadeserv.NewApp(rt.Config{BasePath: "/web"},
				bytes.NewReader(indexBytes), r.URL.Path, httpBkn)
			err := wadeserv.StartRender(app, client.App{app}, w)

			if err != nil {
				if DevMode {
					panic(err)
				} else {
					log.Println(err)
				}
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/web/", http.StatusFound)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
