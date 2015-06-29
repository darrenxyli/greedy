package api

import (
	"github.com/darrenxyli/greedy/database/postgre"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	// DB is global Database for UBDOOR-API
	DB *postgre.ProjectDB
)

func init() {
	DB = postgre.NewProjectDB(
		"127.0.0.1",
		5432,
		"ubdoor",
		"postgres",
		"2jaqx97j")

	DB.CreateProjectTable()
}

// Run is main API
func Run() {
	router := gin.Default()

	// Default Config
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Addtional Settings
	router.GET("/robots.txt", robots)

	v1 := router.Group("/v1")
	{
		// GET information of particular node or all nodes
		v1.GET("/nodes/:name", getNode)
		// POST to create a new node
		v1.POST("/nodes", createNode)
		// GET topic from id
		v1.GET("/topics/:tid", getTopic)
		//POST to create to new topic
		v1.POST("/topics", createTopic)
		// GET reply by id
		v1.GET("/reply/:rid", getReplyByID)
		// GET replies by topic id
		v1.GET("/replies", getReyliesByTopicID)
		// POST to create a new reply
		v1.POST("/reply", createReply)
		// GET to get info of member
		v1.GET("/members", getMember)
		// POST to create a new member
		v1.POST("/members", createMember)
	}

	router.Run(":3448")
}

func robots(c *gin.Context) {
	c.String(
		200,
		"User-agent: *\nDisallow: /\nAllow: /$\nAllow: /debug\nDisallow: /debug/*?taskid=*")
}
