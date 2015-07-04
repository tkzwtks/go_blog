package models

import (
	"time"

	"github.com/tkzwtks/go_photo_blog/db"
)

type Album struct {
	ID        int `json:"-"`
	Title     string
	Photos    []Photo
	CreatedAt time.Time
}

type Photo struct {
	ID        int    `json:"-"`
	AlbumID   int    `json:"-"`
	FilePath  string `sql:"size:256"`
	CreatedAt time.Time
}

func (a Album) FindOne(id int) Album {
	album := Album{}
	conn := db.GetConnection()
	conn.Debug().First(&album, id)
	return album
}

func (a Album) FindAlbumPhotos(albumId int) []Photo {
	var album Album
	var photos []Photo

	conn := db.GetConnection()
	conn.Debug().First(&album, albumId)
	conn.Debug().Model(&album).Related(&photos)

	return photos
}
