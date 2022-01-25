package service

import (
	model "github.com/AndyMile/articles/app/models"
	"github.com/AndyMile/articles/app/repository"
)

type ArticleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		articleRepo: articleRepo,
	}
}

func pagination(page int64) (int, int) {
	var offset int
	const limit int = 20

	if page == 0 || page == 1 {
		offset = 0
	}

	if page > 1 {
		offset = limit * (int(page) - 1)
	}

	return offset, limit
}

func (as *ArticleService) GetAll(page int64) ([]model.Article, error) {
	return as.articleRepo.GetAll(pagination(page))
}

func (as *ArticleService) Get(id int64) (model.Article, error) {
	return as.articleRepo.Get(id)
}

func (as *ArticleService) Create(a model.Article) (model.Article, error) {
	return as.articleRepo.Create(a)
}

func (as *ArticleService) Update(a model.Article) (model.Article, error) {
	return as.articleRepo.Update(a)
}

func (as *ArticleService) Delete(a model.Article) error {
	return as.articleRepo.Delete(a)
}
