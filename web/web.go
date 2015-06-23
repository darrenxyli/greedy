package web

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	// Default Config
	router.LoadHTMLGlob("web/templates/index.html")
	router.LoadHTMLGlob("web/templates/debug.html")
	router.Static("/static", "web/static")
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Index Block
	router.GET("/", index)

	// Debug Block
	router.GET("/debug", debug)

	// Addtional Settings
	router.GET("/robots.txt", robots)

	router.Run(":8080")
}

func robots(c *gin.Context) {
	c.String(
		200,
		"User-agent: *\nDisallow: /\nAllow: /$\nAllow: /debug\nDisallow: /debug/*?taskid=*")
}
