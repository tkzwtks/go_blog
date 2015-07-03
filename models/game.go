package models

import (
	"time"

	"github.com/tkzwtks/go_photo_blog/db"
)

type Game struct {
	ID               int
	MatchName        string
	HomeTeam         string
	Home1stHalfScore int
	Home2ndHalfScore int
	AwayTeam         string
	Away1stHalfScore int
	Away2ndHalfScore int
	Place            string
	GameDay          time.Time
	CreatedAt        time.Time
}

func (g Game) FindOne(id int) Game {
	game := Game{}
	conn := db.GetConnection()
	conn.Debug().First(&game, id)
	return game
}
