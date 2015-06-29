package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/darrenxyli/greedy/api/model"
	"github.com/gin-gonic/gin"
)

func getTopic(c *gin.Context) {
	topicID := c.Param("tid")

	topic := DB.GetTopicByID(topicID)
	topic.Node = DB.GetNodeByID(strconv.FormatInt(topic.NodeID, 10))
	topic.Member = DB.GetMemberByID(strconv.FormatInt(topic.MemberID, 10))
	response := model.Response{Status: "found", Result: topic}
	// Get topic information from database
	if len(topic.Title) != 0 {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusNotAcceptable, model.ItemNotFoundErrorResponse())
	}
}

func createTopic(c *gin.Context) {
	topicTitle := c.PostForm("title")
	topicContent := c.PostForm("content")
	topicMemberID := c.PostForm("author_id")
	topiNodeID := c.PostForm("node_id")

	memberID, _ := strconv.Atoi(topicMemberID)
	nodeID, _ := strconv.Atoi(topiNodeID)

	topic := model.Topic{
		Title:    topicTitle,
		Content:  topicContent,
		MemberID: int64(memberID),
		NodeID:   int64(nodeID),
		Created:  time.Now().Unix(),
	}

	if flag := DB.CreateTopic(topic); flag {
		c.JSON(http.StatusOK, model.CreateSuccessResponse(""))
	} else {
		c.JSON(http.StatusConflict, model.ItemExistErrorResponse(""))
	}
}
