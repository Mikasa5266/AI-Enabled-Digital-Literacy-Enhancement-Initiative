package model

import "time"

type PPT struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Title     string    `json:"title" binding:"required"`
    CoverURL  string    `json:"cover_url"`
    FileURL   string    `json:"file_url"`
    Category  string    `json:"category"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}