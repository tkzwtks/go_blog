package models

import (
	"time"

	"github.com/tkzwtks/go_photo_blog/db"
)

type Album struct {
	ID        int
	Title     string
	CreatedAt time.Time
}

func (a Album) FindOne(id int) Album {
	album := Album{}
	conn := db.GetConnection()
	conn.Debug().First(&album, id)
	return album
}
