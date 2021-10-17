package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	hostFlag := flag.String("host", "localhost", "The host name/IP")
	portFlag := flag.Int("port", 8080, "The port to expose the service from")

	flag.Parse()

	router := setupRouter()

	fmt.Printf("Starting Web Page Analyzer on %s:%d\n", *hostFlag, *portFlag)
	router.Run(fmt.Sprintf("%s:%d", *hostFlag, *portFlag))
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/css", "views/css")
	router.Static("/js", "views/js")
	router.LoadHTMLFiles("views/index.html")

	router.GET("/", getHomePage)
	router.GET("/analyze", getAnalysis)

	return router
}
