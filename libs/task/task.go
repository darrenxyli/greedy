package task

import (
	"time"

	"github.com/darrenxyli/greedy/libs/util"
)

// Task schema
type Task struct {
	//taskId
	ID string `gorm:"primary_key"`
	//project name
	Project string
	//url
	URL string
	//status
	Status uint
	//priority
	Priority uint
	//retries
	Retries uint
	//retried
	Retried uint
	//method
	Method string
	//header
	Header string
	//data
	Data string
	//lastCrawlTime
	LastCrawlTime uint32
	//updateTime
	UpdateTime int64
}

// NewTask the task
func NewTask(oURL string, project string, priority uint, retry uint, method string, header string, data string) *Task {
	return &Task{
		ID:            util.MakeHash(oURL),
		Project:       project,
		URL:           "test",
		Status:        StringToStatus("ACTIVE"),
		Priority:      priority,
		Retries:       retry,
		Retried:       0,
		Method:        method,
		Header:        header,
		Data:          data,
		LastCrawlTime: 1,
		UpdateTime:    time.Now().Unix(),
	}
}

// StatusToString to string
func (task *Task) StatusToString() string {
	switch task.Status {
	case 1:
		return "ACTIVE"
	case 2:
		return "SUCCESS"
	case 3:
		return "FAILED"
	case 4:
		return "BAD"
	default:
		return "UNKNOW"
	}
}

// StringToStatus to status int
func StringToStatus(status string) uint {
	switch status {
	case "ACTIVE":
		return 1
	case "SUCCESS":
		return 2
	case "FAILED":
		return 3
	case "BAD":
		return 4
	default:
		return 5
	}
}
