package models

type CommentRequest struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Comment struct {
	ID       uint   `gorm:"primaryKey"`
	Content  string `gorm:"not null"`
	JokeID   uint
	AuthorID uint
	Author   User
}
