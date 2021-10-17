package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/net/html"
)

func main() {
	router := gin.Default()

	router.Static("/css", "views/css")
	router.Static("/js", "views/js")
	router.LoadHTMLFiles("views/index.html")

	router.GET("/", getHomePage)
	router.GET("/analyze", getAnalysis)

	router.Run("localhost:8080")
}

func getHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
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

	visit(page, &info, resp.Request.URL)
	c.IndentedJSON(http.StatusOK, info)
}
