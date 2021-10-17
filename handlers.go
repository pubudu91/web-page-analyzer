package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

func getHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func getAnalysis(c *gin.Context) {
	url := c.Request.URL.Query().Get("url")
	resp, err := http.Get(url)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var info PageInfo
	info.Status = resp.Status

	page, err := html.Parse(resp.Body)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	analyze(page, &info, resp.Request.URL)
	c.IndentedJSON(http.StatusOK, info)
}
