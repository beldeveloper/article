package controller

import (
	"context"
	"encoding/json"
	model "github.com/AndyMile/articles/app/models"
	pb "github.com/AndyMile/articles/app/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BaseHandler struct {
	grpcClient pb.ArticleClient
}

func NewBaseHandler(client pb.ArticleClient) *BaseHandler {
	return &BaseHandler{
		grpcClient: client,
	}
}

func (h *BaseHandler) GetAll(page int64) ([]model.Article, error) {
	res, err := h.grpcClient.GetAll(context.Background(), &pb.GetAllArticlesRequest{Page: page})
	if err != nil {
		return nil, err
	}
	articles := make([]model.Article, len(res.Articles))
	for i, a := range res.Articles {
		articles[i], err = h.convertArticleBack(a)
		if err != nil {
			return nil, err
		}
	}
	return articles, nil
}

func (h *BaseHandler) Get(articleID int64) (model.Article, error) {
	res, err := h.grpcClient.Get(context.Background(), &pb.GetArticleRequest{Id: articleID})
	if err != nil {
		return model.Article{}, err
	}
	return h.convertArticleBack(res.Article)
}

func (h *BaseHandler) Create(a model.Article) (model.Article, error) {
	pbArticle, err := h.convertArticle(a)
	if err != nil {
		return model.Article{}, err
	}
	res, err := h.grpcClient.Create(context.Background(), &pb.CreateArticleRequest{Article: pbArticle})
	if err != nil {
		return model.Article{}, err
	}
	return h.convertArticleBack(res.Article)
}

func (h *BaseHandler) Update(a model.Article) (model.Article, error) {
	pbArticle, err := h.convertArticle(a)
	if err != nil {
		return model.Article{}, err
	}
	res, err := h.grpcClient.Update(context.Background(), &pb.UpdateArticleRequest{Article: pbArticle})
	if err != nil {
		return model.Article{}, err
	}
	return h.convertArticleBack(res.Article)
}

func (h *BaseHandler) Delete(a model.Article) error {
	pbArticle, err := h.convertArticle(a)
	if err != nil {
		return err
	}
	_, err = h.grpcClient.Delete(context.Background(), &pb.DeleteArticleRequest{Article: pbArticle})
	return err
}

func (h *BaseHandler) convertArticle(article model.Article) (*pb.ArticleItem, error) {
	pbArticle := pb.ArticleItem{
		Id:          article.Id,
		Title:       article.Title,
		Content:     article.Content,
		DateCreated: timestamppb.New(article.DateCreated),
	}
	err := json.Unmarshal(article.Tags, &pbArticle.Tags)
	return &pbArticle, err
}

func (h *BaseHandler) convertArticleBack(pbArticle *pb.ArticleItem) (model.Article, error) {
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
