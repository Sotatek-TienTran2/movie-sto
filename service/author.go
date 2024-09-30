package service

import (
	"movie-sto/dto"
	"movie-sto/errs"
	"movie-sto/model"
	"movie-sto/repository"
)

type AuthorServices interface {
	ListAuthor() (*dto.GetAllAuthorResponse, *errs.AppError)
	CreateAuthor(req dto.CreateAutherRequest) (*dto.CreateAuthorResponse, *errs.AppError)
	SearchAuthor(req dto.SearchAuthorRequest) (*dto.SearchAuthorResponse, *errs.AppError)
	ShowMovieByAuthor(req string) (*dto.ShowMovieByAuthorResponse, *errs.AppError)
}

type DefaultAuthorServices struct {
	repo repository.AuthorRepository
}

func NewAuthorServices(repo repository.AuthorRepository) AuthorServices {

	return DefaultAuthorServices{

		repo: repo,
	}
}

func (a DefaultAuthorServices) ShowMovieByAuthor(req string) (*dto.ShowMovieByAuthorResponse, *errs.AppError) {
	username := req
	res, err := a.repo.ShowMovieByAuthor(username)
	if err != nil {
		return nil, err
	}
	return &dto.ShowMovieByAuthorResponse{Movies: res}, nil
}

func (a DefaultAuthorServices) ListAuthor() (*dto.GetAllAuthorResponse, *errs.AppError) {

	authors, err := a.repo.List()
	if err != nil {

		return nil, err
	}
	return &dto.GetAllAuthorResponse{Authors: authors}, nil
}

func (a DefaultAuthorServices) CreateAuthor(req dto.CreateAutherRequest) (*dto.CreateAuthorResponse, *errs.AppError) {
	author := model.Author{
		Name:       req.Name,
		NativeLand: req.NativeLand,
	}
	newAthor, err := a.repo.Create(author)
	if err != nil {
		return nil, err
	}
	return &dto.CreateAuthorResponse{Author: newAthor}, nil

}

func (a DefaultAuthorServices) SearchAuthor(req dto.SearchAuthorRequest) (*dto.SearchAuthorResponse, *errs.AppError) {
	Name := req.Name
	res, e := a.repo.SearchAuthor(Name)
	if e != nil {

		return nil, e
	}

	return &dto.SearchAuthorResponse{Authors: res}, nil

}
