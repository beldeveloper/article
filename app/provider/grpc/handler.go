package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	model "github.com/AndyMile/articles/app/models"
	pb "github.com/AndyMile/articles/app/proto"
	"github.com/AndyMile/articles/app/service/article"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	as service.Service
	pb.UnimplementedArticleServer
}

func NewHandler(articleService *service.ArticleService) *Handler {
	return &Handler{
		as: articleService,
	}
}

func (h *Handler) GetAll(ctx context.Context, req *pb.GetAllArticlesRequest) (*pb.GetAllArticlesResponse, error) {
	var res pb.GetAllArticlesResponse
	articles, err := h.as.GetAll(req.GetPage())
	if err != nil {
		return &res, err
	}

	res.Articles = make([]*pb.ArticleItem, len(articles))
	for i, v := range articles {
		res.Articles[i], err = h.convertArticle(v)
		if err != nil {
			fmt.Println(err)
		}
	}

	return &res, err
}

func (h *Handler) Get(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	var res pb.GetArticleResponse
	article, err := h.as.Get(req.GetId())
	if err != nil {
		return &res, err
	}
	res.Article, err = h.convertArticle(article)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	var res pb.CreateArticleResponse
	article, err := h.convertArticleBack(req.Article)
	if err != nil {
		return &res, err
	}
	article, err = h.as.Create(article)
	if err != nil {
		return &res, err
	}
	res.Article, err = h.convertArticle(article)
	if err != nil {
		return &res, err
	}
	res.Success = true
	return &res, nil
}

func (h *Handler) Update(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	var res pb.UpdateArticleResponse
	article, err := h.convertArticleBack(req.Article)
	if err != nil {
		return &res, err
	}
	article, err = h.as.Update(article)
	if err != nil {
		return &res, err
	}
	res.Article, err = h.convertArticle(article)
	if err != nil {
		return &res, err
	}
	res.Success = true
	return &res, nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	var res pb.DeleteArticleResponse
	article := model.Article{Id: req.GetArticle().GetId()}
	err := h.as.Delete(article)
	if err != nil {
		return &res, err
	}
	res.Success = true
	return &res, nil
}

func (h *Handler) convertArticle(article model.Article) (*pb.ArticleItem, error) {
	pbArticle := pb.ArticleItem{
		Id:          article.Id,
		Title:       article.Title,
		Content:     article.Content,
		DateCreated: timestamppb.New(article.DateCreated),
	}
	err := json.Unmarshal(article.Tags, &pbArticle.Tags)
	return &pbArticle, err
}

func (h *Handler) convertArticleBack(pbArticle *pb.ArticleItem) (model.Article, error) {
	article := model.Article{
		Id:          pbArticle.Id,
		Title:       pbArticle.Title,
		Content:     pbArticle.Content,
		DateCreated: pbArticle.DateCreated.AsTime(),
	}
	var err error
	article.Tags, err = json.Marshal(pbArticle)
	return article, err
}
