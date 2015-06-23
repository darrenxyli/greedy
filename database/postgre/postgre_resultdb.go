package postgre

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/darrenxyli/greedy/libs/result"
	"github.com/jinzhu/gorm"
)

// ResultDB is the taskdb
type ResultDB struct {
	Projects     []string
	DatabaseName string
	Connection   gorm.DB
}

// NewResultDB create new task database
func NewResultDB(host string, port int, database string, user string, passwd string, projects []string) *ResultDB {
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
		db.Table(GenerateTabelName(name)).CreateTable(&result.Result{})
	}

	return &ResultDB{Projects: projects, DatabaseName: database, Connection: db}
}

// createProjectTable to create TableName
func (resultdb *ResultDB) createProjectTable(project string) {

	tableName := GenerateTabelName(project)

	resultdb.Connection.Table(tableName).CreateTable(&result.Result{})
}

// Insert the data
func (resultdb *ResultDB) Insert(ta *result.Result) bool {
	success := resultdb.Connection.Table(ta.Project).NewRecord(*ta)
	resultdb.Connection.Table(ta.Project).Create(ta)
	return success
}
