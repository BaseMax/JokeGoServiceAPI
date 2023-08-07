package models

type JokeRequest struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Rating  uint   `json:"rating"`
}

type Joke struct {
	ID       uint   `gorm:"primaryKey"`
	Content  string `gorm:"not null"`
	Rating   uint   `gorm:"default=0"`
	AuthorID uint
	Author   User
	Comments []Comment
}
