package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/darrenxyli/greedy/api/model"
	"github.com/gin-gonic/gin"
)

func getHotTopics(c *gin.Context) {
	topics := DB.GetHotTopics()
	var topicsResponse []model.Topic

	for _, topic := range topics {
		topic.Node = DB.GetNodeByID(strconv.FormatInt(topic.NodeID, 10))
		topic.Member = DB.GetMemberByID(strconv.FormatInt(topic.MemberID, 10))
		topicsResponse = append(topicsResponse, topic)
	}
	c.JSON(http.StatusOK, topicsResponse)
}

func getLastedTopics(c *gin.Context) {
	topics := DB.GetLastedTopics()
	var topicsResponse []model.Topic

	for _, topic := range topics {
		topic.Node = DB.GetNodeByID(strconv.FormatInt(topic.NodeID, 10))
		topic.Member = DB.GetMemberByID(strconv.FormatInt(topic.MemberID, 10))
		topicsResponse = append(topicsResponse, topic)
	}
	c.JSON(http.StatusOK, topicsResponse)
}

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
	topicNodeID := c.PostForm("node_id")

	memberID, _ := strconv.Atoi(topicMemberID)
	nodeID, _ := strconv.Atoi(topicNodeID)

	topic := model.Topic{
		Title:          topicTitle,
		Content:        topicContent,
		MemberID:       int64(memberID),
		NodeID:         int64(nodeID),
		Created:        time.Now().Unix(),
		LastedModified: time.Now().Unix(),
	}

	if DB.CreateTopic(topic) && increTopics(topicNodeID) {
		c.JSON(http.StatusOK, model.CreateSuccessResponse(""))
	} else {
		c.JSON(http.StatusConflict, model.ItemExistErrorResponse(""))
	}
}

func increReplies(topicID string) bool {
	DB.IncreaseRepliesCounter(topicID)
	return true
}
