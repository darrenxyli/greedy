package fetcher

import (
	"bytes"

	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"

	"strings"

	"github.com/darrenxyli/greedy/libs/page"
	"github.com/darrenxyli/greedy/libs/request"
	"github.com/darrenxyli/greedy/libs/util"
	"github.com/hu17889/go_spider/core/common/mlog"
	"golang.org/x/net/html/charset"
)

// The HTTPDownloader download page by package net/http.
// The "html" content is contained in dom parser of package goquery.
// The "json" content is saved.
// The "jsonp" content is modified to json.
// The "text" content will save body plain text only.
// The page result is saved in Page.
type HTTPDownloader struct {
}

// NewHTTPDownloader : default constructor
func NewHTTPDownloader() *HTTPDownloader {
	return &HTTPDownloader{}
}

// Download : overload from interface of fetcher.
func (downloader *HTTPDownloader) Download(req *request.Request) *page.Page {

	var mtype string
	var pageItem = page.NewPage(req)
	mtype = req.GetResponceType()

	switch mtype {
	case "html":
		return downloader.downloadHTML(pageItem, req)
	case "json":
		return downloader.downloadJSON(pageItem, req)
	case "jsonp":
		return downloader.downloadJSON(pageItem, req)
	case "text":
		return downloader.downloadText(pageItem, req)
	default:
		mlog.LogInst().LogError("error request type:" + mtype)
	}
	return pageItem
}

// changeCharsetEncodingAuto: Charset auto determine. Use golang.org/x/net/html/charset. Get page body and change it to utf-8
func (downloader *HTTPDownloader) changeCharsetEncodingAuto(contentTypeStr string, sor io.ReadCloser) string {
	var err error
	destReader, err := charset.NewReader(sor, contentTypeStr)

	if err != nil {
		mlog.LogInst().LogError(err.Error())
		destReader = sor
	}

	var sorbody []byte
	if sorbody, err = ioutil.ReadAll(destReader); err != nil {
		mlog.LogInst().LogError(err.Error())
		// For gb2312, an error will be returned.
		// Error like: simplifiedchinese: invalid GBK encoding
		// return ""
	}
	//e,name,certain := charset.DetermineEncoding(sorbody,contentTypeStr)
	bodystr := string(sorbody)

	return bodystr
}

// connectByHTTP : choose http GET/method to download
func connectByHTTP(pageItem *page.Page, req *request.Request) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: req.GetRedirectFunc(),
	}

	httpreq, err := http.NewRequest(req.GetMethod(), req.GetURL(), strings.NewReader(req.GetPostdata()))
	if header := req.GetHeader(); header != nil {
		httpreq.Header = req.GetHeader()
	}

	if cookies := req.GetCookies(); cookies != nil {
		for i := range cookies {
			httpreq.AddCookie(cookies[i])
		}
	}

	var resp *http.Response
	if resp, err = client.Do(httpreq); err != nil {
		if e, ok := err.(*url.Error); ok && e.Err != nil && e.Err.Error() == "normal" {
			//  normal
		} else {
			mlog.LogInst().LogError(err.Error())
			pageItem.SetStatus(true, err.Error())
			return nil, err
		}
	}

	return resp, nil
}

// connectByHTTPProxy : choose a proxy server to excute http GET/method to download
func connectByHTTPProxy(pageItem *page.Page, inReq *request.Request) (*http.Response, error) {
	request, _ := http.NewRequest("GET", inReq.GetURL(), nil)
	proxy, err := url.Parse(inReq.GetProxyHost())
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// downloadFile : Download file and change the charset of page charset.
func (downloader *HTTPDownloader) downloadFile(pageItem *page.Page, req *request.Request) (*page.Page, string) {
	var err error
	var urlstr string
	if urlstr = req.GetURL(); len(urlstr) == 0 {
		mlog.LogInst().LogError("url is empty")
		pageItem.SetStatus(true, "url is empty")
		return pageItem, ""
	}

	var resp *http.Response

	if proxystr := req.GetProxyHost(); len(proxystr) != 0 {
		//using http proxy
		//fmt.Print("HttpProxy Enter ",proxystr,"\n")
		resp, err = connectByHTTPProxy(pageItem, req)
	} else {
		//normal http download
		//fmt.Print("Http Normal Enter \n",proxystr,"\n")
		resp, err = connectByHTTP(pageItem, req)
	}

	if err != nil {
		return pageItem, ""
	}

	//b, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("Resp body %v \r\n", string(b))

	pageItem.SetHeader(resp.Header)
	pageItem.SetCookies(resp.Cookies())

	// get converter to utf-8
	bodyStr := downloader.changeCharsetEncodingAuto(resp.Header.Get("Content-Type"), resp.Body)
	//fmt.Printf("utf-8 body %v \r\n", bodyStr)

	// close response later
	defer resp.Body.Close()

	// return Page, and its body string
	return pageItem, bodyStr
}

// downloadHTML : download html content
func (downloader *HTTPDownloader) downloadHTML(pageItem *page.Page, req *request.Request) *page.Page {
	var err error
	pageItem, destbody := downloader.downloadFile(pageItem, req)
	//fmt.Printf("Destbody %v \r\n", destbody)
	if !pageItem.IsSucc() {
		//fmt.Print("pageItem error \r\n")
		return pageItem
	}
	bodyReader := bytes.NewReader([]byte(destbody))

	// Use goquery parser to test
	var doc *goquery.Document
	if doc, err = goquery.NewDocumentFromReader(bodyReader); err != nil {
		mlog.LogInst().LogError(err.Error())
		pageItem.SetStatus(true, err.Error())
		return pageItem
	}

	var body string
	if body, err = doc.Html(); err != nil {
		mlog.LogInst().LogError(err.Error())
		pageItem.SetStatus(true, err.Error())
		return pageItem
	}

	pageItem.SetBodyStr(body).SetHTMLParser(doc).SetStatus(false, "")

	return pageItem
}

// downloadJSON : download JSON content
func (downloader *HTTPDownloader) downloadJSON(pageItem *page.Page, req *request.Request) *page.Page {
	var err error
	page, destbody := downloader.downloadFile(pageItem, req)
	if !page.IsSucc() {
		return page
	}

	var body []byte
	body = []byte(destbody)
	mtype := req.GetResponceType()
	if mtype == "jsonp" {
		tmpstr := util.JsonpToJson(destbody)
		body = []byte(tmpstr)
	}

	var r *simplejson.Json
	if r, err = simplejson.NewJson(body); err != nil {
		mlog.LogInst().LogError(string(body) + "\t" + err.Error())
		page.SetStatus(true, err.Error())
		return page
	}

	// json result
	page.SetBodyStr(string(body)).SetJSON(r).SetStatus(false, "")

	return page
}

// downloadText : download plain text
func (downloader *HTTPDownloader) downloadText(page *page.Page, req *request.Request) *page.Page {
	page, destbody := downloader.downloadFile(page, req)
	if !page.IsSucc() {
		return page
	}

	page.SetBodyStr(destbody).SetStatus(false, "")
	return page
}
