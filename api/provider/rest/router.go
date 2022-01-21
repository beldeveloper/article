package rest

import (
	controller "github.com/AndyMile/articles/api/controllers"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateRouter(c *controller.BaseHandler) {
	h := NewRouterHandler(c)
	router := gin.Default()

	router.GET("/article/:id", h.Get)
	router.GET("/articles/:page", h.GetAll)
	router.GET("/articles", h.GetAll)
	router.POST("article/create", h.Create)
	router.PUT("article/update", h.Update)
	router.DELETE("article/delete", h.Delete)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
