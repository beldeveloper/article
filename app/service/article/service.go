package service

import (
	model "github.com/AndyMile/articles/app/models"
)

type Service interface {
	GetAll(page int64) ([]model.Article, error)
	Get(id int64) (model.Article, error)
	Create(a model.Article) (model.Article, error)
	Update(a model.Article) (model.Article, error)
	Delete(a model.Article) error
}
