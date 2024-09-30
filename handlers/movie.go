package handlers

import (
	"movie-sto/dto"
	"movie-sto/errs"
	"movie-sto/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHander struct {
	services service.MovieService
}

func NewMovieHandler(services service.MovieService) MovieHander {
	return MovieHander{

		services: services,
	}
}

func (a MovieHander) IndexMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := a.services.IndexMovie()
		if err != nil {

			WriteError(c, err)
			return
		}
		WriteRespon(c, http.StatusOK, res)
	}
}

func (a MovieHander) CreateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie dto.CreateMovieRequest
		err := c.BindJSON(&movie)
		if err != nil {

			WriteError(c, errs.ErrorReadRequestBody())
			return
		}
		_, e := a.services.CreateMovie(movie)
		if e != nil {

			WriteError(c, e)
			return
		}
		WriteRespon(c, http.StatusOK, dto.MessageCreateSuccess("Movie"))
	}
}

func (a MovieHander) SearchMovieByAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var author dto.SearchMovieByAuthorRequest
		err := c.BindJSON(&author)
		if err != nil {

			WriteError(c, errs.ErrorReadRequestBody())
			return
		}
		res, e := a.services.SearchMovieByAuthor(author)
		if e != nil {

			WriteError(c, e)
			return
		}

		WriteRespon(c, http.StatusOK, res)
	}
}

func (a MovieHander) SearchMovieByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category dto.SearchMovieByCategoryRequest
		err := c.BindJSON(&category)
		if err != nil {

			WriteError(c, errs.ErrorReadRequestBody())
			return
		}
		res, e := a.services.SearchMovieByCategory(category)
		if e != nil {

			WriteError(c, e)
			return
		}

		WriteRespon(c, http.StatusOK, res)
	}
}
