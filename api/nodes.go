package api

import (
	"net/http"
	"time"

	"github.com/darrenxyli/greedy/api/model"
	"github.com/gin-gonic/gin"
)

func getNode(c *gin.Context) {
	nodeName := c.Param("name")

	node := DB.GetNodeByName(nodeName)
	response := model.Response{Status: "found", Result: node}
	// Get node information from database
	if len(node.Name) != 0 {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotAcceptable, model.ItemNotFoundErrorResponse())
	}
}

func createNode(c *gin.Context) {
	nodeName := c.PostForm("name")
	nodeTitle := c.PostForm("title")

	node := model.Node{
		Name:    nodeName,
		Title:   nodeTitle,
		Created: time.Now().Unix(),
	}

	if flag, result := DB.CreateNode(node); flag {
		c.JSON(http.StatusOK, model.CreateSuccessResponse(result))
	} else {
		c.JSON(http.StatusConflict, model.ItemExistErrorResponse(result))
	}
}
