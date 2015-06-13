// Package request implements request entity contains url and other relevant informaion.
package request

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bitly/go-simplejson"
)

// Request represents object waiting for being crawled.
type Request struct {
	url string

	// Responce type: html json jsonp text
	respType string

	// GET POST
	method string

	// POST data
	postdata string

	// name for marking url and distinguish different urls in PageProcesser and Pipeline
	urltag string

	// http header
	header http.Header

	// http cookies
	cookies []*http.Cookie

	//proxy host, example='localhost:80'
	proxyHost string

	// Redirect function for downloader used in http.Client
	// If CheckRedirect returns an error, the Client's Get
	// method returns both the previous Response.
	// If CheckRedirect returns error.New("normal"), the error
	// process after client.Do will ignore the error.
	checkRedirect func(req *http.Request, via []*http.Request) error

	meta interface{}
}

// NewRequest returns initialized Request object.
// The respType is json, jsonp, html, text
//
// Example: How to create a Request object
// func NewRequestSimple(url string, respType string, urltag string) *Request {
//     return &Request{url:url, respType:respType}
// }
func NewRequest(url string, respType string, urltag string, method string,
	postdata string, header http.Header, cookies []*http.Cookie,
	checkRedirect func(req *http.Request, via []*http.Request) error,
	meta interface{}) *Request {
	return &Request{url, respType, method, postdata, urltag, header, cookies, "", checkRedirect, meta}
}

// NewRequestWithProxy : create request object with proxy
func NewRequestWithProxy(url string, respType string, urltag string, method string,
	postdata string, header http.Header, cookies []*http.Cookie, proxyHost string,
	checkRedirect func(req *http.Request, via []*http.Request) error,
	meta interface{}) *Request {
	return &Request{url, respType, method, postdata, urltag, header, cookies, proxyHost, checkRedirect, meta}
}

// NewRequestWithHeaderFile : create new request with header
func NewRequestWithHeaderFile(url string, respType string, headerFile string) *Request {
	_, err := os.Stat(headerFile)
	if err != nil {
		//file is not exist , using default mode
		return NewRequest(url, respType, "", "GET", "", nil, nil, nil, nil)
	}

	h := readHeaderFromFile(headerFile)

	return NewRequest(url, respType, "", "GET", "", h, nil, nil, nil)
}

func readHeaderFromFile(headerFile string) http.Header {
	//read file , parse the header and cookies
	b, err := ioutil.ReadFile(headerFile)
	if err != nil {
		//make be:  share access error
		return nil
	}
	js, _ := simplejson.NewJson(b)
	//constructed to header

	h := make(http.Header)
	h.Add("User-Agent", js.Get("User-Agent").MustString())
	h.Add("Referer", js.Get("Referer").MustString())
	h.Add("Cookie", js.Get("Cookie").MustString())
	h.Add("Cache-Control", "max-age=0")
	h.Add("Connection", "keep-alive")
	return h
}

//AddHeaderFile : point to a json file
/* xxx.json
{
	"User-Agent":"curl/7.19.3 (i386-pc-win32) libcurl/7.19.3 OpenSSL/1.0.0d",
	"Referer":"http://weixin.sogou.com/gzh?openid=oIWsFt6Sb7aZmuI98AU7IXlbjJps",
	"Cookie":""
}
*/
func (req *Request) AddHeaderFile(headerFile string) *Request {
	_, err := os.Stat(headerFile)
	if err != nil {
		return req
	}
	h := readHeaderFromFile(headerFile)
	req.header = h
	return req
}

// AddProxyHost : add proxy on request
func (req *Request) AddProxyHost(host string) *Request {
	req.proxyHost = host
	return req
}

// GetURL : get url of request
func (req *Request) GetURL() string {
	return req.url
}

// GetURLTag : get url tag of request
func (req *Request) GetURLTag() string {
	return req.urltag
}

// GetMethod : get HTTP method of request
func (req *Request) GetMethod() string {
	return req.method
}

// GetPostdata : get POST data of request
func (req *Request) GetPostdata() string {
	return req.postdata
}

// GetHeader : get HTTP header of request
func (req *Request) GetHeader() http.Header {
	return req.header
}

// GetCookies : get cookie of request
func (req *Request) GetCookies() []*http.Cookie {
	return req.cookies
}

// GetProxyHost : get proxy of request
func (req *Request) GetProxyHost() string {
	return req.proxyHost
}

// GetResponceType : get reponse type, could be HTML, JSON, etc
func (req *Request) GetResponceType() string {
	return req.respType
}

// GetRedirectFunc : get Redirect functions
func (req *Request) GetRedirectFunc() func(req *http.Request, via []*http.Request) error {
	return req.checkRedirect
}

// GetMeta : get metadata of request
func (req *Request) GetMeta() interface{} {
	return req.meta
}
