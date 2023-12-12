package models

import (
	"gorm.io/gorm"

	"github.com/BaseMax/JokeGoServiceAPI/db"
)

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

func CreateJoke(j *JokeRequest) error {
	var u User
	db := db.GetDB()
	r := db.Find(&u, "username = ?", j.Author)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	joke := Joke{
		Content:  j.Content,
		Rating:   j.Rating,
		AuthorID: u.ID,
	}
	err := db.Create(&joke).Error

	j.ID = joke.ID
	j.Rating = joke.Rating
	return err
}

func FetchAJoke(id uint) (*JokeRequest, error) {
	var j Joke
	r := db.GetDB().Preload("Author").Find(&j, id)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	joke := &JokeRequest{ID: j.ID, Content: j.Content, Author: j.Author.Username, Rating: j.Rating}
	return joke, nil
}

func FetchAllJokes(limit int, page int, sort string) (*[]JokeRequest, uint, error) {
	var count int64
	db := db.GetDB()
	db.Model(&Joke{}).Count(&count)

	var jokes []Joke

	col := "id"
	if sort == "rating" {
		col = "rating"
	}
	r := db.Limit(limit).Offset(((page - 1) * limit)).Order(col + " DESC").Preload("Author").Find(&jokes)
	if r.RowsAffected == 0 {
		return nil, uint(count), gorm.ErrRecordNotFound
	}

	var jokeReqs []JokeRequest
	for _, j := range jokes {
		joke := JokeRequest{ID: j.ID, Content: j.Content,
			Author: j.Author.Username, Rating: j.Rating}
		jokeReqs = append(jokeReqs, joke)
	}
	return &jokeReqs, uint(count), nil
}

func FetchRandomJoke() (*JokeRequest, error) {
	var j Joke
	r := db.GetDB().Order(db.GetRandFunction()).Preload("Author").First(&j)
	if r.Error != nil || r.RowsAffected == 0 {
		return nil, r.Error
	}

	joke := &JokeRequest{ID: j.ID, Content: j.Content,
		Author: j.Author.Username, Rating: j.Rating}
	return joke, nil
}

func FetchTopRatedJokes(limit int) (*[]JokeRequest, error) {
	jokes, _, err := FetchAllJokes(limit, 1, "rating")
	return jokes, err
}

func FetchJokesByAuthor(author string) (*[]JokeRequest, error) {
	var u User
	db := db.GetDB()
	r := db.Preload("Jokes").First(&u, "username = ?", author)
	if r.RowsAffected == 0 {
		return nil, r.Error
	}

	var jokes []JokeRequest
	for _, j := range u.Jokes {
		joke := JokeRequest{ID: j.ID, Content: j.Content,
			Author: author, Rating: j.Rating}
		jokes = append(jokes, joke)
	}
	return &jokes, nil
}

func UpdateJoke(id uint, j *JokeRequest) error {
	var user User
	db := db.GetDB()

	r := db.First(&user, "username = ?", j.Author)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	r = db.Where(id).Updates(Joke{AuthorID: user.ID, Content: j.Content, Rating: j.Rating})
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteJokeById(id uint) error {
	r := db.GetDB().Delete(&Joke{}, id)
	if r.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func RateJoke(id uint, rating uint) (*JokeRequest, error) {
	var joke Joke
	db := db.GetDB()

	r := db.Preload("Author").First(&joke, id)
	if r.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	joke.Rating = rating
	err := db.Save(&joke).Error

	jokeReq := &JokeRequest{ID: id, Content: joke.Content,
		Author: joke.Author.Username, Rating: rating}
	return jokeReq, err
}
