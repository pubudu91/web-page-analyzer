package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/net/html"
)

func main() {
	router := gin.Default()

	router.GET("/analyze", getAnalysis)

	router.Run("localhost:8080")
}

func getAnalysis(c *gin.Context) {
	url := c.Request.URL.Query().Get("url")
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	var info PageInfo
	info.Status = resp.Status

	page, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}

	visit(page, &info, resp.Request.Host)
	c.IndentedJSON(http.StatusOK, info)
}
