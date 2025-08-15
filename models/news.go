package models

import "time"

type News struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Author         string    `json:"author"`
	PublishedAt    time.Time `json:"publishedAt"`
	CommentCount   int       `json:"commentCount"`
	HeaderImageURL string    `json:"headerImageURL"`
}
