package service

import (
	"encoding/json"
	"fmt"
	model "github.com/AndyMile/articles/app/models"
	pb "github.com/AndyMile/articles/app/proto"
	"github.com/AndyMile/articles/app/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(articleRepo *repository.ArticleRepo) *ArticleService {
	return &ArticleService{
		articleRepo: articleRepo,
	
	}
}

func pagination(page int64) (int, int)  {
	var offset int
	const limit int = 20

	if (page == 0 || page == 1) {
		offset = 0
	}

	if (page > 1) {
		offset = limit*(int(page)-1)
	}

	return offset, limit
}

func (as *ArticleService) GetAll(page int64) (*pb.GetAllArticlesResponse, error) {
	
	offset, limit := pagination(page)

	articles, err := as.articleRepo.GetAll(offset, limit)

	items := []*pb.ArticleItem{}

	for _,v := range articles {
		tags := []string{}
		json.Unmarshal([]byte(v.Tags), &tags)
		if err != nil {
			fmt.Println(err)
		}

		a := &pb.ArticleItem {
			Id: int64(v.Id),
			Title: v.Title,
			Content: v.Content,
			DateCreated: timestamppb.New(v.DateCreated),
			Tags: tags,
		}

		items = append(items, a)
	}

	return &pb.GetAllArticlesResponse{Articles: items}, err
}

func (as *ArticleService) Get(id int64) (*pb.GetArticleResponse, error) {
	article, err := as.articleRepo.Get(id)

	tags := []string{}
	json.Unmarshal([]byte(article.Tags), &tags)

	item := pb.ArticleItem {
		Id: int64(article.Id),
		Title: article.Title,
		Content: article.Content,
		DateCreated: timestamppb.New(article.DateCreated),
		Tags: tags,
	}

	return &pb.GetArticleResponse{Article: &item}, err
}

func (as *ArticleService) Create(a model.Article) (*pb.CreateArticleResponse, error) {
	article, err := as.articleRepo.Create(a)
	
	tags := []string{}
	json.Unmarshal([]byte(article.Tags), &tags)

	item := pb.ArticleItem {
		Id: int64(article.Id),
		Title: article.Title,
		Content: article.Content,
	    DateCreated: timestamppb.New(article.DateCreated),
		Tags: tags,
	}

	return &pb.CreateArticleResponse{Article: &item, Success: true}, err
}

func (as *ArticleService) Update(a model.Article) (*pb.UpdateArticleResponse, error) {
	article, err := as.articleRepo.Update(a)

	tags := []string{}
	json.Unmarshal([]byte(article.Tags), &tags)

	item := pb.ArticleItem {
		Id: int64(article.Id),
		Title: article.Title,
		Content: article.Content,
	    DateCreated: timestamppb.New(article.DateCreated),
		Tags: tags,
	}

	return &pb.UpdateArticleResponse{Article: &item, Success: true}, err
}

func (as *ArticleService) Delete(a model.Article) (*pb.DeleteArticleResponse, error) {
	result, err := as.articleRepo.Delete(a)
	return &pb.DeleteArticleResponse{Success: result}, err
}