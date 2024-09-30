package service

import (
	"movie-sto/dto"
	"movie-sto/errs"
	"movie-sto/model"
	"movie-sto/repository"
)

type MovieService interface {
	IndexMovie() (*dto.GetAllMovieResponse, *errs.AppError)
	CreateMovie(req dto.CreateMovieRequest) (*dto.CreateMovieResponse, *errs.AppError)
	SearchMovieByAuthor(req dto.SearchMovieByAuthorRequest) (*dto.SearchMovieByAuthorResponse, *errs.AppError)
	SearchMovieByCategory(req dto.SearchMovieByCategoryRequest) (*dto.SearchMovieByCategoryResponse, *errs.AppError)
}

type DefaultMovieService struct {
	repo repository.MovieRepository
}

func NewMovieServices(repo repository.MovieRepository) MovieService {
	return DefaultMovieService{

		repo: repo,
	}
}

func (a DefaultMovieService) IndexMovie() (*dto.GetAllMovieResponse, *errs.AppError) {
	movies, err := a.repo.IndexMovie()
	if err != nil {

		return nil, err
	}
	return &dto.GetAllMovieResponse{Movies: movies}, nil
}

func (a DefaultMovieService) CreateMovie(req dto.CreateMovieRequest) (*dto.CreateMovieResponse, *errs.AppError) {
	movie := model.Movie{
		Name:         req.Name,
		NameOfAuthor: req.NameOfAuthor,
		Category:     req.Category,
	}
	newMovie, err := a.repo.CreateMovie(movie)
	if err != nil {
		return nil, err
	}
	return &dto.CreateMovieResponse{Movie: newMovie}, nil
}

func (a DefaultMovieService) SearchMovieByAuthor(req dto.SearchMovieByAuthorRequest) (*dto.SearchMovieByAuthorResponse, *errs.AppError) {
	Name := req.NameOfAuthor
	res, err := a.repo.SearchMovieByAuthor(Name)
	if err != nil {
		return nil, err
	}
	return &dto.SearchMovieByAuthorResponse{Movies: res}, nil
}
func (a DefaultMovieService) SearchMovieByCategory(req dto.SearchMovieByCategoryRequest) (*dto.SearchMovieByCategoryResponse, *errs.AppError) {
	Category := req.Category
	res, err := a.repo.SearchMovieByCategory(Category)
	if err != nil {
		return nil, err
	}
	return &dto.SearchMovieByCategoryResponse{Movies: res}, nil
}
