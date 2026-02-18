package model

import "time"

type PPT struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Title          string    `json:"title" form:"title" binding:"required"`
	CoverURL       string    `json:"cover_url"`
	FileURL        string    `json:"file_url"`
	Category       string    `json:"category" form:"category"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Description    string    `json:"description" form:"description"`
	File_size      uint      `json:"file_size"`
	Download_count uint      `json:"download_count"`
}
