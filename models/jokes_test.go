package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/BaseMax/JokeGoServiceAPI/db"
)

func TestCreateJoke(t *testing.T) {
	joke := JokeRequest{Content: "Content 1", Author: FAKE_USER, Rating: 2}
	if assert.NoError(t, CreateJoke(&joke)) {
		assert.NotEmpty(t, joke.ID)
	}

	joke = JokeRequest{Content: "Content 1", Author: "wronguser"}
	if err := CreateJoke(&joke); assert.Error(t, err) {
		assert.Empty(t, joke.ID)
		assert.Equal(t, err, gorm.ErrRecordNotFound)
	}
}

func TestFetchAJoke(t *testing.T) {
	joke, err := FetchAJoke(1)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, joke)
	}

	joke, err = FetchAJoke(1500)
	if assert.Error(t, err) {
		assert.Empty(t, joke)
		assert.Equal(t, err, gorm.ErrRecordNotFound)
	}
}

func TestFetchAllJokes(t *testing.T) {
	var expectedResult []JokeRequest
	for i := 3; i > 1; i-- {
		joke := JokeRequest{ID: uint(i), Content: fmt.Sprint("Content ", i), Author: FAKE_USER, Rating: uint(i * 2)}
		expectedResult = append(expectedResult, joke)
	}

	for i := 2; i <= 5; i++ {
		assert.NoError(t, CreateJoke(&JokeRequest{Content: fmt.Sprint("Content ", i), Author: FAKE_USER, Rating: uint(i * 2)}))
	}

	actualResult, total, err := FetchAllJokes(2, 2, "latest")
	if assert.NoError(t, err) {
		assert.Equal(t, uint(5), total)
		assert.Equal(t, &expectedResult, actualResult)
	}

	actualResult, total, err = FetchAllJokes(2, 2, "rating")
	if assert.NoError(t, err) {
		assert.Equal(t, uint(5), total)
		assert.Equal(t, &expectedResult, actualResult)
	}

	_, total, err = FetchAllJokes(2, 20, "latest")
	if assert.Error(t, err) {
		assert.Equal(t, uint(5), total)
		assert.Equal(t, err, gorm.ErrRecordNotFound)
	}
}

func TestFetchRandomJoke(t *testing.T) {
	joke, err := FetchRandomJoke()
	if assert.NoError(t, err) {
		assert.NotEmpty(t, joke)
	}
}

func TestFetchTopRatedJokes(t *testing.T) {
	var expectedResult []JokeRequest
	for i := 5; i > 2; i-- {
		joke := JokeRequest{ID: uint(i), Content: fmt.Sprint("Content ", i), Author: FAKE_USER, Rating: uint(i * 2)}
		expectedResult = append(expectedResult, joke)
	}

	actualResult, err := FetchTopRatedJokes(3)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedResult, actualResult)
	}
}

func TestFetchJokesByAuthor(t *testing.T) {
	var expectedResult []JokeRequest
	for i := 1; i <= 5; i++ {
		joke := JokeRequest{ID: uint(i), Content: fmt.Sprint("Content ", i), Author: FAKE_USER, Rating: uint(i * 2)}
		expectedResult = append(expectedResult, joke)
	}

	actualResult, err := FetchJokesByAuthor(FAKE_USER)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedResult, actualResult)
	}
}

func TestUpdateJoke(t *testing.T) {
	expectedResult := &JokeRequest{ID: 1, Content: "Updated Content", Author: "newuser", Rating: 100}
	assert.NoError(t, RegisterUser(&User{Username: "newuser", Password: "pass"}))
	assert.NoError(t, UpdateJoke(1, expectedResult))
	actualResult, err := FetchAJoke(1)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedResult, actualResult)
	}

	expectedResult = &JokeRequest{Content: "Updated Content", Author: "wronguser", Rating: 100}
	assert.Error(t, gorm.ErrRecordNotFound, UpdateJoke(1, expectedResult))

	expectedResult = &JokeRequest{Content: "Updated Content", Author: FAKE_USER, Rating: 100}
	assert.Error(t, gorm.ErrRecordNotFound, UpdateJoke(1500, expectedResult))
}

func TestDeleteJokeByIdNotFound(t *testing.T) {
	assert.NoError(t, DeleteJokeById(1))
	assert.NoError(t, DeleteUserByName("newuser"))
	assert.Equal(t, gorm.ErrRecordNotFound, DeleteJokeById(1500))

}

func TestRateJoke(t *testing.T) {
	expectedResult := &JokeRequest{
		ID:      2,
		Content: "Content 2",
		Author:  FAKE_USER,
		Rating:  500,
	}

	actualResult, err := RateJoke(2, 500)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedResult, actualResult)
	}

	_, err = RateJoke(1500, 500)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestTruncateJokes(t *testing.T) {
	db.TruncateTable("comments", "jokes", "users")
}

func TestFetchRandomJokeNotFound(t *testing.T) {
	joke, err := FetchRandomJoke()
	if assert.Error(t, err) {
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Empty(t, joke)
	}
}

func TestFetchTopRatedJokesNotFound(t *testing.T) {
	actualResult, err := FetchTopRatedJokes(3)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Empty(t, actualResult)
}

func TestFetchJokesByAuthorNotFound(t *testing.T) {
	actualResult, err := FetchJokesByAuthor("wronguser")
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Empty(t, actualResult)
}
