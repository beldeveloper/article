package repository

import model "github.com/AndyMile/articles/app/models"

type ArticleRepository interface {
	GetAll(offset int, limit int) ([]model.Article, error)
	Get(id int64) (model.Article, error)
	Create(a model.Article) (model.Article, error)
	Update(a model.Article) (model.Article, error)
	Delete(a model.Article) (bool, error)
}