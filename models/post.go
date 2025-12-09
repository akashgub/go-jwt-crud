package models

type Post struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
	UserID  uint
}
