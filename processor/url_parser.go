package processor

import (
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

type JobResult_t struct {
	Results []UrlResult_t
	Urls    map[string]bool
}

type UrlResult_t struct {
	Url    string
	Images []string
}

func main() {

	crawlUrl("http://tumakan.blog66.fc2.com/")

}

func crawlUrl(urlToCrawl string) {
	var (
		err      error
		content  string
		imgs     []string
		urlToGet *url.URL
		//links    []string
	)

	// Parse URL
	if urlToGet, err = url.Parse(urlToCrawl); err != nil {
		log.Println(err)
		return
	}

	// Retrieve content of URL
	if content, err = getUrlContent(urlToGet.String()); err != nil {
		log.Println(err)
		return
	}

	// Clean up HTML entities
	content = html.UnescapeString(content)

	//if links, err = parseLinks(urlToGet, content); err != nil {
	//	log.Println(err)
	//	return
	//}
	// Retrieve image URLs
	if imgs, err = parseImages(urlToGet, content); err != nil {
		log.Println(err)
		return
	}

	//for _, link := range links {
	//	defer crawlUrl(link)
	//}
	for _, img := range imgs {
		log.Println(img)
	}
} //crawlUrl

func getUrlContent(urlToGet string) (string, error) {
	var (
		err     error
		content []byte
		resp    *http.Response
	)

	// GET content of URL
	if resp, err = http.Get(urlToGet); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if request was successful
	if resp.StatusCode != 200 {
		return "", err
	}

	// Read the body of the HTTP response
	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	}

	return string(content), err
}

func parseImages(urlToGet *url.URL, content string) ([]string, error) {
	var (
		err        error
		imgs       []string
		matches    [][]string
		findImages = regexp.MustCompile("<img.*?src=\"(.*?)\"")
	)

	// Retrieve all image URLs from string
	matches = findImages.FindAllStringSubmatch(content, -1)

	for _, val := range matches {
		var imgUrl *url.URL

		// Parse the image URL
		if imgUrl, err = url.Parse(val[1]); err != nil {
			return imgs, err
		}

		// If the URL is absolute, add it to the slice
		// If the URL is relative, build an absolute URL
		if imgUrl.IsAbs() {
			imgs = append(imgs, imgUrl.String())
		} else {
			imgs = append(imgs, urlToGet.Scheme+"://"+urlToGet.Host+imgUrl.String())
		}
	}

	return imgs, err
}

func parseLinks(urlToGet *url.URL, content string) ([]string, error) {
	var (
		err       error
		links     []string = make([]string, 0)
		matches   [][]string
		findLinks = regexp.MustCompile("<a.*?href=\"(.*?)\"")
	)

	// Retrieve all anchor tag URLs from string
	matches = findLinks.FindAllStringSubmatch(content, -1)

	for _, val := range matches {
		var linkUrl *url.URL

		// Parse the anchr tag URL
		if linkUrl, err = url.Parse(val[1]); err != nil {
			return links, err
		}

		// If the URL is absolute, add it to the slice
		// If the URL is relative, build an absolute URL
		if linkUrl.IsAbs() {
			links = append(links, linkUrl.String())
		} else {
			links = append(links, urlToGet.Scheme+"://"+urlToGet.Host+linkUrl.String())
		}
	}

	return links, err
}
