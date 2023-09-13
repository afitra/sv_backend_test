package models

import (
	"time"
)

type Post struct {
	Id          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Content     string    `db:"content" json:"content"`
	Category    string    `db:"category" json:"category"`
	CreatedDate time.Time `db:"created_date" json:"created_date"`
	UpdatedDate time.Time `db:"updated_date" json:"updated_date"`
	Status      string    `db:"status" json:"status"`
}

type PayloadPost struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

type ArticleCreated struct {
	Count int    `json:"count"`
	Post  []Post `json:"post"`
}
