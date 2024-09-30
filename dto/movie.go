package dto

import "movie-sto/model"

type GetAllMovieResponse struct {
	Movies []model.Movie
}

type CreateMovieRequest struct {
	Name         string `json:"Name,omitempty"`
	NameOfAuthor string `json:"NameOfAuthor,omitempty"`
	Category     string `json:"Category,omitempty"`
}

type CreateMovieResponse struct {
	Movie model.Movie
}

type SearchMovieByCategoryRequest struct {
	Category string `json:"Category"`
}

type SearchMovieByCategoryResponse struct {
	Movies []model.Movie
}

type SearchMovieByAuthorRequest struct {
	NameOfAuthor string `json:"NameOfAuthor,omitempty"`
}

type SearchMovieByAuthorResponse struct {
	Movies []model.Movie
}
