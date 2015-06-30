package api

import (
	"net/http"
	"time"

	"github.com/darrenxyli/greedy/api/model"
	"github.com/gin-gonic/gin"
)

func getMember(c *gin.Context) {
	username := c.Query("username")
	userID := c.Query("id")

	if len(username) != 0 {
		member := DB.GetMemberByName(username)
		response := model.Response{Status: "found", Result: member}
		// Get node information from database
		if len(member.Username) != 0 {
			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(http.StatusNotAcceptable, model.ItemNotFoundErrorResponse())
		}
	} else if len(userID) != 0 {
		member := DB.GetMemberByID(userID)
		response := model.Response{Status: "found", Result: member}
		// Get node information from database
		if len(member.Username) != 0 {
			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(http.StatusNotAcceptable, model.ItemNotFoundErrorResponse())
		}
	} else {
		c.JSON(http.StatusBadRequest, model.BadParametersErrorResponse())
	}
}

func createMember(c *gin.Context) {
	var member model.Member
	c.Bind(&member)
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	// fbID := c.PostForm("fb_id")
	member.Created = time.Now().Unix()

	if flag := DB.CreateMember(member); flag {
		c.JSON(http.StatusOK, model.CreateSuccessResponse(""))
	} else {
		c.JSON(http.StatusConflict, model.ItemExistErrorResponse(""))
	}

	// if len(member.FacebookID) != 0 {
	// 	name := c.PostForm("name")
	// 	fbLink := c.PostForm("fb_link_url")
	// 	miniAvatar := c.PostForm("mini_url")
	// 	normalAvatar := c.PostForm("normal_url")
	// 	largeAvatar := c.PostForm("large_url")
	// 	fbToken := c.PostForm("fb_token")
	//
	// 	member := model.Member{
	// 		Name:          name,
	// 		FacebookID:    fbID,
	// 		Facebook:      fbLink,
	// 		AvatarMini:    miniAvatar,
	// 		AvatarNormal:  normalAvatar,
	// 		AvatarLarge:   largeAvatar,
	// 		FacebookToken: fbToken,
	// 		Created:       time.Now().Unix(),
	// 	}
	//
	// 	if flag := DB.CreateMember(member); flag {
	// 		c.JSON(http.StatusOK, model.CreateSuccessResponse(""))
	// 	} else {
	// 		c.JSON(http.StatusConflict, model.ItemExistErrorResponse(""))
	// 	}
	// } else if len(username) == 0 || len(password) == 0 {
	// 	c.JSON(http.StatusBadRequest, model.BadParametersErrorResponse())
	// } else {
	// 	member := model.Member{
	// 		Username: username,
	// 		Password: password,
	// 		Created:  time.Now().Unix(),
	// 	}
	//
	// 	if flag := DB.CreateMember(member); flag {
	// 		c.JSON(http.StatusOK, model.CreateSuccessResponse(""))
	// 	} else {
	// 		c.JSON(http.StatusConflict, model.ItemExistErrorResponse(""))
	// 	}
	// }
}
