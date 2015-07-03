package models

import (
	"time"

	"github.com/tkzwtks/go_photo_blog/db"
)

type Photo struct {
	ID        int
	FilePath  string `sql:"size:256"`
	AlbumId   int
	CreatedAt time.Time
}

func (p Photo) FindOne(id int) Photo {
	photo := Photo{}
	conn := db.GetConnection()
	conn.Debug().First(&photo, id)
	return photo
}
