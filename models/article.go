package models

import (
	"time"

	"github.com/tkzwtks/go_photo_blog/db"
)

type Article struct {
	ID        int
	Title     string `sql:"size:128"`
	Text      string `sql:"DEFAULT'';type:text"`
	Author    string `sql:"size 20"`
	AlbumID   int
	GameID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a Article) FindOne(id int) Article {
	article := Article{}
	conn := db.GetConnection()
	conn.Debug().First(&article, id)
	return article
}
