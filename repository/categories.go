package repository

import (
	"database/sql"
	"movie-sto/errs"
	"movie-sto/model"

	"github.com/go-redis/redis"
)

type CategoryRepository interface {
	List() ([]model.Categories, *errs.AppError)
	Create(author model.Categories) (model.Categories, *errs.AppError)
	Search(req string) ([]model.Categories, *errs.AppError)
}

type DefaultCategoryRepository struct {
	db    *sql.DB
	redis *redis.Client
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return DefaultCategoryRepository{
		db: db,
	}
}

func (a DefaultCategoryRepository) List() ([]model.Categories, *errs.AppError) {
	var categories []model.Categories
	var category model.Categories
	res, err := a.db.Query("SELECT * FROM dungphung.categories ")
	if err != nil {

		return nil, errs.ErrorGetData()
	}
	for res.Next() {
		err = res.Scan(&category.IdCategories, &category.Category)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		categories = append(categories, category)
	}
	return categories, nil
}

func (a DefaultCategoryRepository) Create(category model.Categories) (model.Categories, *errs.AppError) {
	newCategory, err := a.db.Prepare("INSERT INTO dungphung.categories(Category) VALUES (?) ")
	if err != nil {

		return category, errs.ErrorInsertData()
	}
	newCategory.Exec(category.Category)
	return category, nil
}

func (a DefaultCategoryRepository) Search(req string) ([]model.Categories, *errs.AppError) {
	var categories []model.Categories
	var category model.Categories
	if req == "" {
		return categories, nil
	}
	bodyString := "%" + req + "%"
	res, err := a.db.Query("SELECT categories.idCategories,categories.Category FROM dungphung.categories as categories WHERE categories.Category LIKE ?", bodyString)
	if err != nil {
		return nil, errs.ErrorGetData()
	}
	for res.Next() {
		err := res.Scan(&category.IdCategories, &category.Category)
		if err != nil {

			return nil, errs.ErrorReadData()
		}
		categories = append(categories, category)
	}
	return categories, nil
}
