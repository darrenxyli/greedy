package main

import (
	"github.com/darrenxyli/greedy/fetcher"
	"github.com/darrenxyli/greedy/utils/config"
)

func main() {
	config.New()
	fetcher.CrawlURL("http://www.4porn.com")
}
