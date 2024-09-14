package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"

	"github.com/BatmiBoom/web_crawler_go/cmd/crawler"
)

func main() {
	arg_len := len(os.Args[1:])
	if arg_len < 1 {
		fmt.Println("too few arguments")
		fmt.Println("Usage: crawler <url> <maxpages> <maxconcurrency>")
		os.Exit(1)
	} else if arg_len > 3 {
		fmt.Println("too many arguments provided")
		fmt.Println("Usage: crawler <url> <maxpages> <maxconcurrency>")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	maxPages, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("The second argument has to be a number")
	}
	maxConcurrency, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("The third argument has to be a number")
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("ERROR: Parsing base url %v\n", err)
	}

	fmt.Println("Generating Config")
	cfg := crawler.Config{
		Pages:              map[string]int{},
		MaxPages:           maxPages,
		MaxConcurrency:     maxConcurrency,
		BaseURL:            parsedBaseURL,
		Mu:                 &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}),
		Wg:                 &sync.WaitGroup{},
	}

	fmt.Printf("Starting crawl of: %s\n", baseURL)
	pages, err := cfg.CrawlPage(cfg.BaseURL)
	if err != nil {
		fmt.Printf("There was an error in the crawl %v\n", err)
	}

	fmt.Println(pages)
}
