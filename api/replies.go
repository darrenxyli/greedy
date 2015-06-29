package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/darrenxyli/greedy/api/model"
	"github.com/gin-gonic/gin"
)

// getReplyByID to get a single reply by reply ID
func getReplyByID(c *gin.Context) {
	replyID := c.Param("rid")

	if len(replyID) == 0 {
		c.JSON(http.StatusBadRequest, model.BadParametersErrorResponse())
	} else {
		reply := DB.GetReplyByID(replyID)
		reply.Member = DB.GetMemberByID(strconv.FormatInt(reply.MemberID, 10))
		response := model.Response{Status: "found", Result: reply}
		// Get topic information from database
		if len(reply.Content) != 0 {
			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(http.StatusNotAcceptable, model.ItemNotFoundErrorResponse())
		}
	}
}

// getReyliesByTopicID to get replies by topic ID
func getReyliesByTopicID(c *gin.Context) {
	topicID := c.Query("topic_id")

	if len(topicID) == 0 {
		c.JSON(http.StatusBadRequest, model.BadParametersErrorResponse())
	} else {
		var repliesResponse []model.Reply
		replies := DB.GetReyliesByTopicID(topicID)
		for _, reply := range replies {
			reply := DB.GetReplyByID(strconv.FormatInt(reply.ID, 10))
			reply.Member = DB.GetMemberByID(strconv.FormatInt(reply.MemberID, 10))
			repliesResponse = append(repliesResponse, reply)
		}
		response := model.Response{Status: "ok", Result: repliesResponse}

		c.JSON(http.StatusOK, response)
	}
}

// createReply to create reply for a topic
func createReply(c *gin.Context) {
	replyContent := c.PostForm("content")
	replyMemberID := c.PostForm("author_id")
	replyTopicID := c.PostForm("topic_id")

	if len(replyContent) == 0 || len(replyMemberID) == 0 || len(replyTopicID) == 0 {
		c.JSON(http.StatusBadRequest, model.BadParametersErrorResponse())
	} else {
		memberID, _ := strconv.Atoi(replyMemberID)
		topicID, _ := strconv.Atoi(replyTopicID)

		reply := model.Reply{
			Content:  replyContent,
			MemberID: int64(memberID),
			TopicID:  int64(topicID),
			Created:  time.Now().Unix(),
		}

		if flag := DB.CreateReply(reply); flag {
			c.JSON(http.StatusOK, model.CreateSuccessResponse(""))
		} else {
			c.JSON(http.StatusConflict, model.ItemExistErrorResponse(""))
		}
	}
}
