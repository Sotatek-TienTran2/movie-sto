package repository

import (
	"database/sql"
	"movie-sto/errs"
	"movie-sto/model"
	"strings"
)

type MovieRepository interface {
	IndexMovie() ([]model.Movie, *errs.AppError)
	CreateMovie(movie model.Movie) (model.Movie, *errs.AppError)
	SearchMovieByAuthor(req string) ([]model.Movie, *errs.AppError)
	SearchMovieByCategory(req string) ([]model.Movie, *errs.AppError)
}

type DefaultMovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {

	return DefaultMovieRepository{

		db: db,
	}
}

func (a DefaultMovieRepository) IndexMovie() ([]model.Movie, *errs.AppError) {
	var movies []model.Movie
	var movie model.Movie
	res, err := a.db.Query("SELECT movie.idMovie,movie.Name FROM dungphung.movie as movie")
	if err != nil {

		return nil, errs.ErrorGetData()
	}
	for res.Next() {
		err = res.Scan(&movie.IdMovie, &movie.Name)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		movies = append(movies, movie)
	}
	return movies, nil
}

func (a DefaultMovieRepository) CreateMovie(movie model.Movie) (model.Movie, *errs.AppError) {
	var idMovie int
	NewMovie, err := a.db.Prepare("INSERT INTO dungphung.movie(Name) VALUES (?) ")
	if err != nil {
		return movie, errs.ErrorInsertData()
	}
	NewMovie.Exec(movie.Name)

	err1 := a.db.QueryRow("SELECT movie.idMovie FROM dungphung.movie as movie WHERE movie.Name = ?", movie.Name).Scan(&idMovie)
	if err1 != nil || err1 == sql.ErrNoRows {
		return movie, errs.ErrorGetData()
	}

	author := strings.Split(movie.NameOfAuthor, "; ")
	for i := 0; i < len(author); i++ {
		var idAuthor int
		err := a.db.QueryRow("SELECT author.idAuthor FROM dungphung.author as author WHERE author.Name = ?", author[i]).Scan(&idAuthor)
		if err != nil || err == sql.ErrNoRows {
			return movie, errs.ErrorGetData()
		}

		NewMovieAuthor, err := a.db.Prepare("INSERT INTO dungphung.movie_author (idAuthor, idMovie) VALUES (?, ?) ")
		if err != nil {
			return movie, errs.ErrorInsertData()
		}
		NewMovieAuthor.Exec(idAuthor, idMovie)
	}

	categories := strings.Split(movie.Category, "; ")
	for i := 0; i < len(categories); i++ {
		var idCategories int
		err1 := a.db.QueryRow("SELECT categories.idCategories FROM dungphung.categories as categories WHERE categories.Category = ?", categories[i]).Scan(&idCategories)
		if err1 != nil || err1 == sql.ErrNoRows {
			return movie, errs.ErrorGetData()
		}

		selDB1, err := a.db.Prepare("INSERT INTO dungphung.movie_categories (idCategories, idMovie) VALUES (?, ?) ")
		if err != nil {
			return movie, errs.ErrorInsertData()
		}
		selDB1.Exec(idCategories, idMovie)
	}
	defer a.db.Close()

	return movie, nil
}

func (a DefaultMovieRepository) SearchMovieByAuthor(req string) ([]model.Movie, *errs.AppError) {
	var movies []model.Movie
	if req == "" {
		return movies, nil
	}
	bodyString := "%" + req + "%"
	res, err := a.db.Query("SELECT movie.Name, author.Name FROM dungphung.movie as movie,  dungphung.author as author, dungphung.movie_author as ba WHERE movie.idMovie = ba.idMovie AND ba.idAuthor = author.idAuthor AND author.Name LIKE ?", bodyString)
	if err != nil {
		return movies, errs.ErrorGetData()
	}
	for res.Next() {
		var movie model.Movie
		err = res.Scan(&movie.Name, &movie.NameOfAuthor)
		if err != nil {
			return nil, errs.ErrorReadData()
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (a DefaultMovieRepository) SearchMovieByCategory(req string) ([]model.Movie, *errs.AppError) {
	var movies []model.Movie
	if req == "" {
		return movies, nil
	}
	bodyString := "%" + req + "%"
	res, err := a.db.Query("SELECT movie.Name, categories.Category FROM dungphung.movie as movie,  dungphung.categories as categories, dungphung.movie_categories as bc WHERE movie.idMovie = bc.idMovie AND bc.idCategories = categories.idCategories AND categories.Category LIKE ?", bodyString)

	if err != nil {
		return movies, errs.ErrorGetData()
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
