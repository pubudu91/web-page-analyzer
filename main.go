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

	router := gin.Default()

	router.Static("/css", "views/css")
	router.Static("/js", "views/js")
	router.LoadHTMLFiles("views/index.html")

	router.GET("/", getHomePage)
	router.GET("/analyze", getAnalysis)

	router.Run(fmt.Sprintf("%s:%d", *hostFlag, *portFlag))
}
