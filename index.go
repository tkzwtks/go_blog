package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/tkzwtks/go_photo_blog/models"
)

type Message struct {
	ID   int64
	Data string
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	mes := Message{1, c.URLParams["name"]}
	b, error := json.Marshal(mes)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Fprintf(w, string(b))
}

func init() {
	fmt.Println("initialized app for tkzwtks\n")
}

func entry(c web.C, w http.ResponseWriter, r *http.Request) {
	// TODO: error
	id, _ := strconv.Atoi(c.URLParams["id"])
	entry := models.FindOneArticle(id)

	b, _ := json.Marshal(entry)

	fmt.Fprintf(w, string(b))
}

func main() {
	goji.Get("/hello/:name", hello)
	goji.Get("/entry/:id", entry)
	goji.Serve()
}
