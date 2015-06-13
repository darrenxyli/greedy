// Package page contains result catched by Downloader.
// And it alse has result parsed by PageProcesser.
package page

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"github.com/darrenxyli/greedy/libs/page_items"
	"github.com/darrenxyli/greedy/libs/request"
	"github.com/hu17889/go_spider/core/common/mlog"
)

// Page represents an entity be crawled.
type Page struct {
	// The isfail is true when crawl process is failed and errormsg is the fail resean.
	isfail   bool
	errormsg string

	// The request is crawled by spider that contains url and relevent information.
	req *request.Request

	// The body is plain text of crawl result.
	body string

	header  http.Header
	cookies []*http.Cookie

	// The docParser is a pointer of goquery boject that contains html result.
	docParser *goquery.Document

	// The jsonMap is the json result.
	jsonMap *simplejson.Json

	// The pItems is object for save Key-Values in PageProcesser.
	// And pItems is output in Pipline.
	pItems *page_items.PageItems

	// The targetRequests is requests to put into Scheduler.
	targetRequests []*request.Request
}

// NewPage returns initialized Page object.
func NewPage(req *request.Request) *Page {
	return &Page{pItems: page_items.NewPageItems(req), req: req}
}

// SetHeader save the header of http responce
func (p *Page) SetHeader(header http.Header) {
	p.header = header
}

// GetHeader returns the header of http responce
func (p *Page) GetHeader() http.Header {
	return p.header
}

// SetCookies save the cookies of http responce
func (p *Page) SetCookies(cookies []*http.Cookie) {
	p.cookies = cookies
}

// GetCookies returns the cookies of http responce
func (p *Page) GetCookies() []*http.Cookie {
	return p.cookies
}

// IsSucc test whether download process success or not.
func (p *Page) IsSucc() bool {
	return !p.isfail
}

// Errormsg show the download error message.
func (p *Page) Errormsg() string {
	return p.errormsg
}

// SetStatus save status info about download process.
func (p *Page) SetStatus(isfail bool, errormsg string) {
	p.isfail = isfail
	p.errormsg = errormsg
}

// AddField saves KV string pair to PageItems preparing for Pipeline
func (p *Page) AddField(key string, value string) {
	p.pItems.AddItem(key, value)
}

// GetPageItems returns PageItems object that record KV pair parsed in PageProcesser.
func (p *Page) GetPageItems() *page_items.PageItems {
	return p.pItems
}

// SetSkip set label "skip" of PageItems.
// PageItems will not be saved in Pipeline wher skip is set true
func (p *Page) SetSkip(skip bool) {
	p.pItems.SetSkip(skip)
}

// GetSkip returns skip label of PageItems.
func (p *Page) GetSkip() bool {
	return p.pItems.GetSkip()
}

// SetRequest saves request oject of p page.
func (p *Page) SetRequest(r *request.Request) *Page {
	p.req = r
	return p
}

// GetRequest returns request oject of p page.
func (p *Page) GetRequest() *request.Request {
	return p.req
}

// GetURLTag returns name of url.
func (p *Page) GetURLTag() string {
	return p.req.GetURLTag()
}

// AddTargetRequest adds one new Request waitting for crawl.
func (p *Page) AddTargetRequest(url string, respType string) *Page {
	p.targetRequests = append(p.targetRequests, request.NewRequest(url, respType, "", "GET", "", nil, nil, nil, nil))
	return p
}

// AddTargetRequests adds new Requests waitting for crawl.
func (p *Page) AddTargetRequests(urls []string, respType string) *Page {
	for _, url := range urls {
		p.AddTargetRequest(url, respType)
	}
	return p
}

// AddTargetRequestWithProxy adds one new Request waitting for crawl.
func (p *Page) AddTargetRequestWithProxy(url string, respType string, proxyHost string) *Page {

	p.targetRequests = append(p.targetRequests, request.NewRequestWithProxy(url, respType, "", "GET", "", nil, nil, proxyHost, nil, nil))
	return p
}

// AddTargetRequestsWithProxy adds new Requests waitting for crawl.
func (p *Page) AddTargetRequestsWithProxy(urls []string, respType string, proxyHost string) *Page {
	for _, url := range urls {
		p.AddTargetRequestWithProxy(url, respType, proxyHost)
	}
	return p
}

// AddTargetRequestWithHeaderFile adds one new Request with header file for waitting for crawl.
func (p *Page) AddTargetRequestWithHeaderFile(url string, respType string, headerFile string) *Page {
	p.targetRequests = append(p.targetRequests, request.NewRequestWithHeaderFile(url, respType, headerFile))
	return p
}

// AddTargetRequestWithParams adds one new Request waitting for crawl.
// The respType is "html" or "json" or "jsonp" or "text".
// The urltag is name for marking url and distinguish different urls in PageProcesser and Pipeline.
// The method is POST or GET.
// The postdata is http body string.
// The header is http header.
// The cookies is http cookies.
func (p *Page) AddTargetRequestWithParams(req *request.Request) *Page {
	p.targetRequests = append(p.targetRequests, req)
	return p
}

// AddTargetRequestsWithParams adds new Requests waitting for crawl.
func (p *Page) AddTargetRequestsWithParams(reqs []*request.Request) *Page {
	for _, req := range reqs {
		p.AddTargetRequestWithParams(req)
	}
	return p
}

// GetTargetRequests returns the target requests that will put into Scheduler
func (p *Page) GetTargetRequests() []*request.Request {
	return p.targetRequests
}

// SetBodyStr saves plain string crawled in Page.
func (p *Page) SetBodyStr(body string) *Page {
	p.body = body
	return p
}

// GetBodyStr returns plain string crawled.
func (p *Page) GetBodyStr() string {
	return p.body
}

// SetHTMLParser saves goquery object binded to target crawl result.
func (p *Page) SetHTMLParser(doc *goquery.Document) *Page {
	p.docParser = doc
	return p
}

// GetHTMLParser returns goquery object binded to target crawl result.
func (p *Page) GetHTMLParser() *goquery.Document {
	return p.docParser
}

// ReSetHTMLParser returns goquery object binded to target crawl result.
func (p *Page) ReSetHTMLParser() *goquery.Document {
	r := strings.NewReader(p.body)
	var err error
	p.docParser, err = goquery.NewDocumentFromReader(r)
	if err != nil {
		mlog.LogInst().LogError(err.Error())
		panic(err.Error())
	}
	return p.docParser
}

// SetJSON saves json result.
func (p *Page) SetJSON(js *simplejson.Json) *Page {
	p.jsonMap = js
	return p
}

// GetJSON returns json result.
func (p *Page) GetJSON() *simplejson.Json {
	return p.jsonMap
}
