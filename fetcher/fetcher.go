package fetcher

import (
	"github.com/darrenxyli/greedy/libs/page"
	"github.com/darrenxyli/greedy/libs/request"
)

// The Downloader interface.
// You can implement the interface by implement function Download.
// Function Download need to return Page instance pointer that has request result downloaded from Request.
type Downloader interface {
	Download(req *request.Request) *page.Page
}
