package models

import (
	"fmt"
	"testing"

	"github.com/BaseMax/JokeGoServiceAPI/db"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateComment(t *testing.T) {
	joke := &JokeRequest{Content: "Joke", Author: FAKE_USER, Rating: 15}
	assert.NoError(t, CreateJoke(joke))
	assert.NoError(t, CreateComment(1, &CommentRequest{Content: "Comment 1", Author: FAKE_USER}))
	assert.Equal(t, gorm.ErrRecordNotFound, CreateComment(1, &CommentRequest{Content: "Comment 1", Author: "wronguser"}))
	assert.Error(t, CreateComment(1500, &CommentRequest{Content: "Comment 1", Author: FAKE_USER}))
}

func TestFetchCommentById(t *testing.T) {
	expectedResult := &CommentRequest{ID: 1, Content: "Comment 1", Author: FAKE_USER}
	actualResult, err := FetchCommentById(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, actualResult)

	actualResult, err = FetchCommentById(1500)
	assert.Error(t, gorm.ErrRecordNotFound, err)
	assert.Empty(t, actualResult)
}

func TestFetchAllComments(t *testing.T) {
	expectedResult := []CommentRequest{{ID: 1, Content: "Comment 1", Author: FAKE_USER}}
	for i := 2; i <= 5; i++ {
		comment := CommentRequest{ID: uint(i), Content: fmt.Sprint("Comment ", i), Author: FAKE_USER}
		expectedResult = append(expectedResult, comment)
		assert.NoError(t, CreateComment(1, &comment))
	}

	actualResult, err := FetchAllComments(1)
	assert.NoError(t, err)
	assert.Equal(t, &expectedResult, actualResult)
}

func TestUpdateComment(t *testing.T) {
	assert.NoError(t, UpdateComment(1, &CommentRequest{Content: "Updated Comment", Author: FAKE_USER}))
	assert.Error(t, UpdateComment(1500, &CommentRequest{Content: "Updated Comment", Author: FAKE_USER}))
}

func TestDeleteComment(t *testing.T) {
	assert.NoError(t, DeleteComment(1))
	assert.Equal(t, gorm.ErrRecordNotFound, DeleteComment(1))
}

func TestTruncateComments(t *testing.T) {
	db := db.GetDB()

	db.Raw("DELETE FROM comments;").Row()
	db.Raw("ALTER TABLE comments AUTO_INCREMENT=1;").Row()

	db.Raw("DELETE FROM jokes;").Row()
	db.Raw("ALTER TABLE jokes AUTO_INCREMENT=1;").Row()
}

func TestFetchAllCommentsNotFound(t *testing.T) {
	_, err := FetchAllComments(1)
	assert.Error(t, gorm.ErrRecordNotFound, err)
}
