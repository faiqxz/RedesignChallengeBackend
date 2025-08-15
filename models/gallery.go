package models

import "time"

type Gallery struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ImageTitle  string    `json:"imageTitle"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageURL"`
	UploadedAt  time.Time `json:"uploadedAt"`
}
