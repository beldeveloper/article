package rest

import (
	controller "github.com/AndyMile/articles/api/controllers"
	"encoding/json"
	"net/http"
	"strconv"
	pb "github.com/AndyMile/articles/app/proto"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type ArticleInput struct {
	Id int `json:"id"`
	Title  string `json:"title"`
	Content string `json:"content"`
	Tags datatypes.JSON `json:"tags"`
}

type routerHandler struct {
	c *controller.BaseHandler
}

func NewRouterHandler(c *controller.BaseHandler) *routerHandler {
	return &routerHandler{
		c: c,
	}
}

func (h routerHandler) GetAll(ctx *gin.Context) {
	var page int64 = 0
	var err error
	p := ctx.Param("page")
	
	if (p != "") {
		page, err = strconv.ParseInt(p, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	res, err := h.c.GetAll(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h routerHandler) Get(ctx *gin.Context) {
	articelId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.c.Get(articelId)
		
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} 
		
	ctx.JSON(http.StatusOK, gin.H{"response": res})	
}

func (h routerHandler) Create(ctx *gin.Context) {
	var input ArticleInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if (input.Title == "") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	tags := []string{}
	err :=json.Unmarshal([]byte(input.Tags), &tags)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	article := pb.ArticleItem{Title: input.Title, Content: input.Content, Tags: tags}

	res, err := h.c.Create(article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

    ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h routerHandler) Update(ctx *gin.Context) {
	var input ArticleInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tags := []string{}
	err := json.Unmarshal([]byte(input.Tags), &tags)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	
	article := pb.ArticleItem{Id: int64(input.Id), Title: input.Title, Content: input.Content, Tags: tags}

	res, err := h.c.Update(article)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    ctx.JSON(http.StatusOK, gin.H{"response": res})
}

func (h routerHandler) Delete(ctx *gin.Context) {
	var input ArticleInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := pb.ArticleItem{Id: int64(input.Id)}

	res, err := h.c.Delete(article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": res})
}