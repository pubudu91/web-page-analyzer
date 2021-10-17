package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAnalysis(t *testing.T) {
	router := setupRouter()
	server := httptest.NewServer(setupMockHandler("test.html"))
	defer server.Close()

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/analyze?url=%s/test", server.URL), nil)
	router.ServeHTTP(writer, req)

	var respJson PageInfo
	json.NewDecoder(writer.Body).Decode(&respJson)

	assert.Equal(t, 200, writer.Code)
	assert.Equal(t, "200 OK", respJson.Status)
	assert.Equal(t, "XHTML 1.0 Strict", respJson.HtmlVersion)
	assert.Equal(t, "Test HTML File", respJson.Title)
	assert.Equal(t, 1, respJson.Headings.H1)
	assert.Equal(t, 1, respJson.Headings.H2)
	assert.Equal(t, 1, respJson.Headings.H3)
	assert.Equal(t, 1, respJson.Headings.H4)
	assert.Equal(t, 1, respJson.Headings.H5)
	assert.Equal(t, 1, respJson.Headings.H6)
	assert.Equal(t, 1, respJson.Links.Internal)
	assert.Equal(t, 1, respJson.Links.External)
	assert.Equal(t, 2, respJson.Links.Inaccessible)
	assert.Equal(t, false, respJson.HasLoginForm)
}

func TestInvalidURL(t *testing.T) {
	router := setupRouter()
	server := httptest.NewServer(setupMockHandler("test.html"))
	defer server.Close()

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/analyze?url=http://foo", nil)
	router.ServeHTTP(writer, req)

	var respJson struct {
		Error string `json:"error"`
	}

	json.NewDecoder(writer.Body).Decode(&respJson)

	assert.Equal(t, 400, writer.Code)
	assert.Equal(t, "Get \"http://foo\": dial tcp: lookup foo: no such host", respJson.Error)
}

func setupMockHandler(file string) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("test-resources/" + file)
	router.GET("/test", func(c *gin.Context) { c.HTML(http.StatusOK, file, nil) })
	return router
}
