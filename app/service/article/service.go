package service

import (
	model "github.com/AndyMile/articles/app/models"
	pb "github.com/AndyMile/articles/app/proto"
)

type Service interface {
	GetAll(page int64) (*pb.GetAllArticlesResponse, error)
	Get(id int64) (*pb.GetArticleResponse, error)
	Create(a model.Article) (*pb.CreateArticleResponse, error)
	Update(a model.Article) (*pb.UpdateArticleResponse, error)
	Delete(a model.Article) (*pb.DeleteArticleResponse, error)
}