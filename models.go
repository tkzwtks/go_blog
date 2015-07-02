package main

import (
	"time"
)

type User struct {
	ID int
	Birthday time.Time
	Age int
	Name string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Article struct {
	ID int
	Title string `sql:"size:128"`
	Text string `sql:"DEFAULT'';type:text"`
	Author string `sql:"size 20"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Photo struct {
	ID int
	FilePath string `sql:"size:256"`
	CreatedAt time.Time
}
