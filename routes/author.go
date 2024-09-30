package routes

import (
	"movie-sto/config"
	"movie-sto/handlers"
	"movie-sto/middlewares"
	"movie-sto/redis"
	"movie-sto/repository"
	"movie-sto/service"

	"github.com/gin-gonic/gin"
)

func AuthorRoute(router *gin.Engine) {

	handler := handlers.NewAuthorHandler(service.NewAuthorServices(repository.NewAuthorRepository(config.DB)))
	jwtMiddleware := middlewares.NewJWTMiddleware(repository.NewAuthorRepository(config.DB), redis.RDB)
	route := router.Group("/api/author")
	{

		route.GET("/", handler.GetListAuthor())
		route.POST("/create", handler.CreateAuthor())
		route.POST("/search", handler.SearchAuthor())
		route.GET("/show", jwtMiddleware.Verify(), handler.ShowMovieByAuthor())
	}
}
