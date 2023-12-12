package models

import (
	"gorm.io/gorm"

	"github.com/BaseMax/JokeGoServiceAPI/db"
)

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

func CreateComment(joke_id uint, c *CommentRequest) error {
	db := db.GetDB()
	var user User
	r := db.Find(&user, "username = ?", c.Author)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	comment := Comment{ID: c.ID, Content: c.Content, JokeID: joke_id, AuthorID: user.ID}
	err := db.Create(&comment).Error

	c.ID = comment.ID
	return err
}

func FetchCommentById(id uint) (*CommentRequest, error) {
	var comment Comment
	db := db.GetDB()

	r := db.Preload("Author").First(&comment, id)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	commentReq := CommentRequest{ID: comment.ID,
		Content: comment.Content, Author: comment.Author.Username}
	return &commentReq, nil
}

func FetchAllComments(joke_id uint) (*[]CommentRequest, error) {
	var joke Joke
	db := db.GetDB()

	r := db.Preload("Comments").Preload("Author").First(&joke, joke_id)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var comments []CommentRequest
	for _, c := range joke.Comments {
		comment := CommentRequest{ID: c.ID, Content: c.Content, Author: joke.Author.Username}
		comments = append(comments, comment)
	}
	return &comments, nil
}

func UpdateComment(id uint, c *CommentRequest) error {
	var comment Comment
	db := db.GetDB()

	r := db.Preload("Author").First(&comment, "id = ?", id)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	err := db.Where(id).Updates(Comment{Content: c.Content}).Error
	c.ID = comment.ID
	c.Author = comment.Author.Username
	return err
}

func DeleteComment(id uint) error {
	r := db.GetDB().Delete(&Comment{}, id)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return r.Error
}
