package controller

import (
	"encoding/json"

	"github.com/tkzwtks/go_photo_blog/models"
)

type ArticleController struct {
}

func (a *ArticleController) Show(id int) string {
	var article *models.Article = &models.Article{}
	entry := article.FindOne(id)
	b, _ := json.Marshal(entry)
	return string(b)
}
