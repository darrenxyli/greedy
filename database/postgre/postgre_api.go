package postgre

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/darrenxyli/greedy/api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// ProjectDB is the taskdb
type ProjectDB struct {
	DatabaseName string
	Connection   gorm.DB
}

// NewProjectDB create new task database
func NewProjectDB(host string, port int, database string, user string, passwd string) *ProjectDB {
	protocal := "postgres://"
	db, err := gorm.Open(
		"postgres",
		strings.Join([]string{
			protocal,
			user, ":", passwd, "@", host, ":", strconv.Itoa(port),
			"/", database, "?", "sslmode=disable"}, ""))

	if err != nil {
		fmt.Println("Can not connect")
		fmt.Println(err)
		os.Exit(1)
		return &ProjectDB{DatabaseName: database}
	}
	// Get the Underlying native Golang database
	// connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	// Then can invoke `*sql.DB`'s functions with it such as
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)

	return &ProjectDB{DatabaseName: database, Connection: db}
}

// CreateProjectTable to create TableName
func (projectdb *ProjectDB) CreateProjectTable() {
	projectdb.Connection.AutoMigrate(
		&model.Node{}, &model.Member{}, &model.Reply{}, &model.Topic{})
}

// Insert the data
func (projectdb *ProjectDB) Insert() bool {
	// success := projectdb.Connection.Table(ta.Project).NewRecord(*ta)
	// projectdb.Connection.Table(ta.Project).Create(ta)
	return true
}

// GetHotTopics get recent hot topics, in recent 1 week
func (projectdb *ProjectDB) GetHotTopics() []model.Topic {
	var topics []model.Topic
	projectdb.Connection.Where("lasted_modified > ?", time.Now().Unix()-604796).Order("replies desc").Limit(20).Find(&topics)
	return topics
}

// GetLastedTopics get recent hot topics
func (projectdb *ProjectDB) GetLastedTopics() []model.Topic {
	var topics []model.Topic
	projectdb.Connection.Order("lasted_modified desc").Limit(20).Find(&topics)
	return topics
}

// GetAllNodes returns all nodes
func (projectdb *ProjectDB) GetAllNodes() []model.Node {
	var nodes []model.Node
	projectdb.Connection.Find(&nodes)
	return nodes
}

// GetNodeByName get node information
func (projectdb *ProjectDB) GetNodeByName(name string) model.Node {
	node := model.Node{}
	projectdb.Connection.Where("name = ?", name).Find(&node)
	return node
}

// GetNodeByID get node information from id
func (projectdb *ProjectDB) GetNodeByID(ID string) model.Node {
	node := model.Node{}
	NodeID, _ := strconv.Atoi(ID)
	projectdb.Connection.Where("id = ?", int64(NodeID)).Find(&node)
	return node
}

// CreateNode create node
func (projectdb *ProjectDB) CreateNode(node model.Node) (bool, model.Node) {

	var flag bool

	existed := projectdb.GetNodeByName(node.Name)

	if flag = len(existed.Name) == 0; flag {
		projectdb.Connection.Create(&node)
		existed = projectdb.GetNodeByName(node.Name)
	}

	return flag, existed
}

// GetMemberByName get member by name
func (projectdb *ProjectDB) GetMemberByName(username string) model.Member {
	member := model.Member{}
	projectdb.Connection.Where("username = ?", username).Find(&member)
	return member
}

// GetMemberByID get member by name
func (projectdb *ProjectDB) GetMemberByID(ID string) model.Member {
	member := model.Member{}
	id, _ := strconv.Atoi(ID)
	projectdb.Connection.Where("id = ?", int64(id)).Find(&member)
	return member
}

// CreateMember create member
func (projectdb *ProjectDB) CreateMember(member model.Member) bool {
	var flag bool

	existed := projectdb.GetMemberByName(member.Username)

	if flag = len(existed.Username) == 0; flag {
		projectdb.Connection.Create(&member)
	}

	return flag
}

// GetTopicByID get node information from id
func (projectdb *ProjectDB) GetTopicByID(ID string) model.Topic {
	topic := model.Topic{}
	TopicID, _ := strconv.Atoi(ID)
	projectdb.Connection.Where("id = ?", int64(TopicID)).Find(&topic)
	return topic
}

// CreateTopic create topic
func (projectdb *ProjectDB) CreateTopic(topic model.Topic) bool {
	projectdb.Connection.Create(&topic)
	return true
}

// CreateReply creates reply
func (projectdb *ProjectDB) CreateReply(reply model.Reply) bool {
	projectdb.Connection.Create(&reply)
	return true
}

// GetReplyByID get reply by id
func (projectdb *ProjectDB) GetReplyByID(ID string) model.Reply {
	reply := model.Reply{}
	replyID, _ := strconv.Atoi(ID)
	projectdb.Connection.Where("id = ?", int64(replyID)).Find(&reply)
	return reply
}

// GetReyliesByTopicID get replies
func (projectdb *ProjectDB) GetReyliesByTopicID(ID string) []model.Reply {
	var replies []model.Reply
	topicID, _ := strconv.Atoi(ID)
	projectdb.Connection.Where("topic_id = ?", int64(topicID)).Find(&replies)
	return replies
}

// IncreaseRepliesCounter increase replies
func (projectdb *ProjectDB) IncreaseRepliesCounter(ID string) {
	topicID, _ := strconv.Atoi(ID)
	projectdb.Connection.Table("topics").Where("id = ?", int64(topicID)).Updates(map[string]interface{}{
		"replies":         gorm.Expr("replies + ?", 1),
		"lasted_modified": time.Now().Unix()})
}

// IncreaseTopicsCounter increase topics
func (projectdb *ProjectDB) IncreaseTopicsCounter(ID string) {
	nodeID, _ := strconv.Atoi(ID)
	projectdb.Connection.Table("nodes").Where("id = ?", int64(nodeID)).Updates(map[string]interface{}{
		"topics": gorm.Expr("topics + ?", 1)})
}
