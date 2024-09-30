package routes

import (
	"movie-sto/config"
	"movie-sto/handlers"
	"movie-sto/repository"
	"movie-sto/service"

	"github.com/gin-gonic/gin"
)

func MovieRoute(router *gin.Engine) {

	handler := handlers.NewMovieHandler(service.NewMovieServices(repository.NewMovieRepository(config.DB)))
	route := router.Group("/api/movie")
	{
		route.GET("/", handler.IndexMovie())
		route.POST("/create", handler.CreateMovie())
		route.POST("/search-by-author", handler.SearchMovieByAuthor())
		route.POST("/search-by-category", handler.SearchMovieByCategory())
	}
}
