package grpc

import (
	"context"
	pb "github.com/AndyMile/articles/app/proto"
	"encoding/json"
	"time"
	model "github.com/AndyMile/articles/app/models"
	"github.com/AndyMile/articles/app/service/article"
)

type Handler struct {
	as *service.ArticleService
	pb.UnimplementedArticleServer
}

func NewHandler(articleService *service.ArticleService) *Handler {
	return &Handler{
		as: articleService,
	}
}

func (h *Handler) GetAll(ctx context.Context, req *pb.GetAllArticlesRequest) (*pb.GetAllArticlesResponse, error) {
	return h.as.GetAll(req.GetPage())
}

func (h *Handler) Get(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	return h.as.Get(req.GetId())
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	article := model.Article{
		Id: req.GetArticle().GetId(), 
		Title: req.GetArticle().GetTitle(),
		Content: req.GetArticle().GetContent(),
		DateCreated: time.Now(),
	}

	if (len(req.GetArticle().GetTags()) > 0) {
		tags, err := json.Marshal(req.GetArticle().GetTags())
		if (err != nil) {
			panic(err)
		}
		article.Tags = tags
	}

	return h.as.Create(article)
}

func (h *Handler) Update(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	article := model.Article{
		Id: req.GetArticle().GetId(), 
		Title: req.GetArticle().GetTitle(),
		Content: req.GetArticle().GetContent(),
		DateCreated: time.Now(),
	}

	if (len(req.GetArticle().GetTags()) > 0) {
		tags, err := json.Marshal(req.GetArticle().GetTags())
		if (err != nil) {
			panic(err)
		}
		article.Tags = tags
	}

	return h.as.Update(article)
}

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	article := model.Article{Id: req.GetArticle().GetId()}
	return h.as.Delete(article)
}