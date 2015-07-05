package main

import (
	"net/http"
	"regexp"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	"github.com/tkzwtks/go_photo_blog/controller"
)

func main() {
	api := web.New()
	goji.Handle("/api/*", api)
	api.Use(middleware.SubRouter)
	api.Get("/entry/:id", controller.ShowArticle)
	api.Get("/album/:id", controller.ShowAlbum)

	goji.Handle("/", http.FileServer(http.Dir("./static")))
	goji.Handle(regexp.MustCompile("^/(css|js|html)"), http.FileServer(http.Dir("./static")))

	goji.Serve()
}
