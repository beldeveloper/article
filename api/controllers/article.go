package controller

import (
	"context"
	pb "github.com/AndyMile/articles/app/proto"
)

type BaseHandler struct {
	grpcClient pb.ArticleClient
}

func NewBaseHandler(client pb.ArticleClient) *BaseHandler {
	return &BaseHandler{
		grpcClient: client,
	}
}

func (h *BaseHandler) GetAll(page int64) (*pb.GetAllArticlesResponse, error) {
	return h.grpcClient.GetAll(context.Background(), &pb.GetAllArticlesRequest{Page: page})	
}

func (h *BaseHandler) Get(articelId int64) (*pb.GetArticleResponse, error) {
	return h.grpcClient.Get(context.Background(), &pb.GetArticleRequest{Id: articelId})
}

func (h *BaseHandler) Create(a pb.ArticleItem) (*pb.CreateArticleResponse, error) {
	return h.grpcClient.Create(context.Background(), &pb.CreateArticleRequest{Article: &a})
}

func (h *BaseHandler) Update(a pb.ArticleItem) (*pb.UpdateArticleResponse, error) {
	return h.grpcClient.Update(context.Background(), &pb.UpdateArticleRequest{Article: &a})
}

func (h *BaseHandler) Delete(a pb.ArticleItem) (*pb.DeleteArticleResponse, error) {
	return h.grpcClient.Delete(context.Background(), &pb.DeleteArticleRequest{Article: &a})
}


