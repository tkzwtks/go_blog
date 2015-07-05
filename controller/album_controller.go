package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tkzwtks/go_photo_blog/models"
	"github.com/zenazn/goji/web"
)

func ShowAlbum(c web.C, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(c.URLParams["id"])
	var album *models.Album = &models.Album{}
	var photos []models.Photo
	photos = album.FindAlbumPhotos(id)

	b, _ := json.Marshal(photos)
	fmt.Fprintf(w, string(b))
}
