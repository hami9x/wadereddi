package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/phaikawl/wade"
	wadeserv "github.com/phaikawl/wade/rbackend/serverside"
	"github.com/phaikawl/wadereddi/client"
)

const (
	ServersidePrerender = false
	DevMode             = true
)

func main() {
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
			wadeserv.RenderApp(w, wade.AppConfig{
				StartPage: "pg-home",
				BasePath:  "/web",
			}, client.InitFunc, bytes.NewReader(indexBytes), http.DefaultServeMux, r, "/api")
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/web/", http.StatusFound)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
