package repository

import (
	"fmt"

	model "github.com/AndyMile/articles/app/models"
	"gorm.io/gorm"
)

var articles []model.Article

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{
		db: db,
	}
}

func (r *ArticleRepo) GetAll(offset int, limit int) ([]model.Article, error) {
	err := r.db.Limit(limit).Offset(offset).Find(&articles).Error
	return articles, err
}

func (r *ArticleRepo) Get(id int64) (model.Article, error) {
	a := model.Article{Id: id}
	err := r.db.First(&a).Error
	return a, err
}

func (r *ArticleRepo) Create(a model.Article) (model.Article, error) {
	err := r.db.Create(&a).Error
	return a, err
}

func (r *ArticleRepo) Update(a model.Article) (model.Article, error) {	
	if err := r.db.First(&a).Error; err != nil {
		return a, err
	}

	err := r.db.Model(&a).Updates(&a).Error	
	fmt.Println(err)
	return a, err
}

func (r *ArticleRepo) Delete(a model.Article) (bool, error) {
	err := r.db.Delete(&a).Error
	result := true

	if (err != nil) {
		result = false
	}

	return result, err
}