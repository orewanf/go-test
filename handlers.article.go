package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context)  {
	articles := getAllArticles()

	// Call the HTMl method of the context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Home Page",
			"payload": articles,
		},
	)
}

func getArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("article_id"))
	var articleTitle = "Not found"
	if err != nil {
		log.Fatal("not found")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	article, err := getArticleByID(articleID)
	if err != nil {
		if err = c.AbortWithError(http.StatusNotFound, err); err != nil {
			log.Fatal(err)
		}
		return
	}

	if article != nil {
		articleTitle = article.Title
	}
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the article.html template
		"article.html",
		// Pass the data that the page uses
		gin.H{
			"title":   articleTitle,
			"payload": article,
		},
	)
}
