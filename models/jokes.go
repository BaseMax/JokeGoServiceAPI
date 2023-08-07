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

func CreateJoke(j *Joke) error {
	return nil
}

func FetchAJoke(id uint) (*Joke, error) {
	return nil, nil
}

func FetchAllJokes(limit uint, page uint, sort string) (*[]Joke, error) {
	return nil, nil
}

func FetchRandomJoke() (*Joke, error) {
	return nil, nil
}

func FetchTopRatedJokes() (*[]Joke, error) {
	return nil, nil
}

func GetJokeByAuthor(author string) (*Joke, error) {
	return nil, nil
}

func UpdateJoke(id uint, j *Joke) error {
	return nil
}
func DeleteJokeById(id uint) error {
	return nil
}
func RateJoke(id uint, rating uint) (*Joke, error) {
	return nil, nil
}
