package main

import (
	"net/http"
	"regexp"

	"github.com/flosch/pongo2"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	"github.com/tkzwtks/go_photo_blog/controller"
	"github.com/tkzwtks/go_photo_blog/models"
)

func main() {
	// setup
	models.Migrate()
	pongo2.DefaultSet.SetBaseDirectory("view")

	api := web.New()
	goji.Handle("/api/*", api)
	api.Use(middleware.SubRouter)
	api.Get("/entry/:id", controller.ShowArticle)
	api.Get("/album/:id", controller.ShowAlbum)

	goji.Handle("/", http.FileServer(http.Dir("./static")))
	goji.Handle(regexp.MustCompile("^/(css|js)"), http.FileServer(http.Dir("./static")))

	goji.Serve()
}
