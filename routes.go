package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Handle the index route
	router.GET("/", showIndexPage)

	// Handle get requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)
}
