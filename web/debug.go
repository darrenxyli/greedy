package web

import (
	"github.com/gin-gonic/gin"
)

func debug(c *gin.Context) {
	c.HTML(200, "debug.html", gin.H{"project_name": "Home"})
}
