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

func CreateComment(c *Comment) error {
	return nil
}

func FetchCommentById(id uint) (*Comment, error) {
	return nil, nil
}

func FetchjAllComments(limit uint, page uint) (*[]Comment, error) {
	return nil, nil
}

func UpdateComment(id uint, c *Comment) (*Comment, error) {
	return nil, nil
}

func DeleteComment(id uint) error {
	return nil
}
