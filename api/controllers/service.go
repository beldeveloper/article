package controller

import pb "github.com/AndyMile/articles/app/proto"

type ArticleApiController interface {
	GetAll(page int64) (*pb.GetAllArticlesResponse, error)
	Get(articelId int64) (*pb.GetArticleResponse, error)
	Create(a *pb.ArticleItem) (*pb.CreateArticleResponse, error)
	Update(a *pb.ArticleItem) (*pb.UpdateArticleResponse, error)
	Delete(a *pb.ArticleItem) (*pb.DeleteArticleResponse, error)
}