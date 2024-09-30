package repository

import (
	"database/sql"
	"movie-sto/errs"
	"movie-sto/model"
)

type AuthorRepository interface {
	List() ([]model.Author, *errs.AppError)
	Create(model.Author) (model.Author, *errs.AppError)
	SearchAuthor(req string) ([]model.Author, *errs.AppError)
	FindAuthorByUsername(username string) (string, *errs.AppError)
	ShowMovieByAuthor(req string) ([]model.Movie, *errs.AppError)
}

type DefaultAuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {

	return DefaultAuthorRepository{

		db: db,
	}
}

func (a DefaultAuthorRepository) FindAuthorByUsername(username string) (string, *errs.AppError) {
	res, err := a.db.Query("select author.username from dungphung.author as author where author.username = ? ", username)
	if err != nil {

		return "false", errs.ErrorGetData()
	}
	var user string
	for res.Next() {
		err = res.Scan(&user)
		if err != nil {

			return "false", errs.ErrorReadData()
		}
	}
	if user == "" {
		return "false", nil
	} else {
		return user, nil
	}
}

func (a DefaultAuthorRepository) List() ([]model.Author, *errs.AppError) {

	res, err := a.db.Query("SELECT author.idauthor, author.Name, author.NativeLand FROM dungphung.author as author")

	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var authors []model.Author
	var author model.Author
	for res.Next() {
		err = res.Scan(&author.IdAuthor, &author.Name, &author.NativeLand)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a DefaultAuthorRepository) Create(author model.Author) (model.Author, *errs.AppError) {

	newAuthor, err := a.db.Prepare("INSERT INTO dungphung.author(Name, NativeLand) VALUES (? ,?) ")
	if err != nil {

		return author, errs.ErrorInsertData()
	}
	newAuthor.Exec(author.Name, author.NativeLand)
	return author, nil
}

func (a DefaultAuthorRepository) SearchAuthor(req string) ([]model.Author, *errs.AppError) {
	var authors []model.Author
	if req == "" {
		return authors, nil
	}
	bodyString := "%" + req + "%"
	res, err := a.db.Query("SELECT author.Name, author.NativeLand FROM dungphung.author as author WHERE author.Name LIKE ?", bodyString)

	if err != nil {

		return nil, errs.ErrorGetData()
	}
	for res.Next() {

		var author model.Author
		err = res.Scan(&author.Name, &author.NativeLand)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a DefaultAuthorRepository) ShowMovieByAuthor(req string) ([]model.Movie, *errs.AppError) {
	var movies = []model.Movie{}
	res, err := a.db.Query("select movie.idMovie, movie.Name from dungphung.movie as movie , dungphung.author as author, dungphung.movie_author as b_a where movie.idMovie = b_a.idMovie and b_a.idAuthor = author.idAuthor and author.username = ?", req)
	if err != nil {

		return nil, errs.ErrorGetData()
	}
	for res.Next() {
		var movie model.Movie
		err = res.Scan(&movie.Name, &movie.Category)
		if err != nil {
			return nil, errs.ErrorReadData()
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
