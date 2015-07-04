package main

import (
	"github.com/tkzwtks/go_photo_blog/models"
	"github.com/zenazn/goji"
)

func main() {
	models.Migrate()
	goji.Get("/entry/:id", showArticle)
	goji.Get("/album/:id", showAlbum)
	goji.Serve()
}
