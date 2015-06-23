package postgre

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/darrenxyli/greedy/libs/task"
	"github.com/jinzhu/gorm"
)

// TaskDB is the taskdb
type TaskDB struct {
	Projects     []string
	DatabaseName string
	Connection   gorm.DB
}

// NewTaskDB create new task database
func NewTaskDB(host string, port int, database string, user string, passwd string, projects []string) *TaskDB {
	protocal := "postgres://"
	// s := "postgres://darrenxyli:2jaqx97j@amazon.cbtwp3cmfmsx.us-west-2.rds.amazonaws.com:5432/ocean?sslmode=disable"
	db, err := gorm.Open(
		"postgres",
		strings.Join([]string{
			protocal,
			user, ":", passwd, "@", host, ":", strconv.Itoa(port),
			"/", database, "?", "sslmode=disable"}, ""))

	if err != nil {
		fmt.Println("Can not connect")
	}

	// Get the Underlying native Golang database
	// connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	// Then can invoke `*sql.DB`'s functions with it such as
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	for _, name := range projects {
		db.Table(GenerateTabelName(name)).CreateTable(&task.Task{})
	}

	return &TaskDB{Projects: projects, DatabaseName: database, Connection: db}
}

// createProjectTable to create TableName
func (taskdb *TaskDB) createProjectTable(project string) {

	tableName := GenerateTabelName(project)

	taskdb.Connection.Table(tableName).CreateTable(&task.Task{})
}

// Insert the data
func (taskdb *TaskDB) Insert(ta *task.Task) bool {
	success := taskdb.Connection.Table(ta.Project).NewRecord(*ta)
	taskdb.Connection.Table(ta.Project).Create(ta)
	return success
}

// Get the data
func (taskdb *TaskDB) Get(project string, ta *task.Task) {
	taskdb.Connection.Table(GenerateTabelName(project)).First(ta)
}

// LoadTasks load the tasks
func (taskdb *TaskDB) LoadTasks(status string, project string, limit int) {
	// rows, err := taskdb.Connection.Table(GenerateTabelName(project)).Where("status = ?", task.StringToStatus(status)).Rows()
	//
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// defer rows.Close()
	//
	// for rows.Next() {
	// 	fmt.Println("OOOOOOOOOO")
	// 	// fmt.Println(rows)
	// 	taskItem := new(task.Task)
	// 	rows.Scan(taskItem.ID, taskItem.Project, taskItem.URL, taskItem.Status,
	// 		taskItem.Priority, taskItem.Retries, taskItem.Retried, taskItem.Method,
	// 		taskItem.Header, taskItem.Data, taskItem.LastCrawlTime, taskItem.UpdateTime)
	//
	// 	fmt.Println(taskItem.Project)
	// }

	objecs := new([]task.Task)

	taskdb.Connection.Table(GenerateTabelName(project)).Where("status = ?", task.StringToStatus(status)).Limit(limit).Find(objecs)

	for _, re := range *objecs {
		fmt.Println(re.Project)
	}

}
