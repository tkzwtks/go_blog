package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tkzwtks/go_photo_blog/controller"
	"github.com/tkzwtks/go_photo_blog/models"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func entry(c web.C, w http.ResponseWriter, r *http.Request) {
	// TODO: error
	id, _ := strconv.Atoi(c.URLParams["id"])
	articleController := &controller.ArticleController{}

	b := articleController.Show(id)
	fmt.Fprintf(w, b)
}

func main() {
	models.Migrate()
	goji.Get("/entry/:id", entry)
	goji.Serve()
}
