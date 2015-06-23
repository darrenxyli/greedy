package result

import (
	"time"

	"github.com/darrenxyli/greedy/libs/util"
)

// Result schema
type Result struct {
	//taskId
	ID string `gorm:"primary_key"`
	//project name
	Project string
	//url
	URL string
	//status
	Duration uint
	//priority
	Site string
	//retries
	Img string
	//retried
	Title string
	//method
	LastCrawlTime int64
}

// NewResult to new the result
func NewResult(oURL string, project string, duration uint, site string, img string, title string) *Result {
	return &Result{
		ID:            util.MakeHash(oURL),
		Project:       project,
		URL:           oURL,
		Duration:      duration,
		Site:          site,
		Img:           img,
		Title:         title,
		LastCrawlTime: time.Now().Unix(),
	}
}
