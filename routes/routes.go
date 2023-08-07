package routes

import (
	"fmt"
	"os"

	"github.com/BaseMax/JokeGoServiceAPI/controllers"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	jwtKey := []byte(os.Getenv("JWT_KET"))
	jwtConf := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		SigningKey: jwtKey,
	}
	g := e.Group("/", echojwt.WithConfig(jwtConf))

	fmt.Println(string(jwtKey))
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	g.POST("refresh", controllers.Refresh)

	g.POST("jokes", controllers.CreateJoke)
	g.GET("jokes/:joke_id", controllers.GetJoke)
	g.GET("jokes", controllers.GetAllJokes)
	g.GET("jokes/random", controllers.GetRandomJoke)
	g.GET("jokes/top-rated", controllers.GetTopRatedJoke)
	g.GET("jokes/authors/:author_name", controllers.GetJokeByAuthor)
	g.PUT("jokes/:joke_id", controllers.EditJoke)
	g.DELETE("jokes/:joke_id", controllers.DeleteJoke)
	g.POST("jokes/:joke_id/rating", controllers.RateJoke)

	g.POST("jokes/:joke_id/comments", controllers.CreateJokeComment)
	g.GET("jokes/:joke_id/comments/:comment_id", controllers.GetJokeComment)
	g.GET("jokes/:joke_id/comments", controllers.GetJokeComments)
	g.PUT("jokes/:joke_id/comments/:comment_id", controllers.EditJokeComment)
	g.DELETE("jokes/:joke_id/comments/:comment_id", controllers.DeleteJokeComment)

	return e
}
