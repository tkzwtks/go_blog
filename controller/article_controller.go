package controller

import (
	"net/http"
	"strconv"

	"github.com/flosch/pongo2"
	"github.com/tkzwtks/go_photo_blog/models"
	"github.com/zenazn/goji/web"
)

func ShowArticle(c web.C, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(c.URLParams["id"])
	var article *models.Article = &models.Article{}
	entry := article.FindOne(id)

	tpl, err := pongo2.DefaultSet.FromFile("article.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteWriter(pongo2.Context{"article": entry}, w)
	//b, _ := json.Marshal(article)
	//	fmt.Fprintf(w, string(b))
}
