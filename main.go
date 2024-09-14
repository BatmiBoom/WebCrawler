package main

import (
	"fmt"
	"os"

	"github.com/BatmiBoom/web_crawler_go/cmd/crawler"
)

func main() {
	arg_len := len(os.Args[1:])
	if arg_len < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if arg_len > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	BASE_URL := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", BASE_URL)

	pages, err := crawler.CrawlPage(BASE_URL, BASE_URL, map[string]int{})
	if err != nil {
		fmt.Printf("There was an error in the crawl %v", err)
	}

	fmt.Println(pages)
}
