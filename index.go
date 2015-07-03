package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tkzwtks/go_photo_blog/models"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func entry(c web.C, w http.ResponseWriter, r *http.Request) {
	// TODO: error
	id, _ := strconv.Atoi(c.URLParams["id"])
	var article *models.Article = &models.Article{}
	entry := article.FindOne(id)
	b, _ := json.Marshal(entry)

	fmt.Fprintf(w, string(b))
}

func main() {
	models.Migrate()
	goji.Get("/entry/:id", entry)
	goji.Serve()
}
