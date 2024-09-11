package main

import (
	"fmt"
	"os"

	handlehtml "github.com/BatmiBoom/web_crawler_go/cmd/handle_html"
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
	fmt.Printf("starting crawl of: %s", BASE_URL)

	rawHTML, err := handlehtml.GetHTML(BASE_URL)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	fmt.Println(rawHTML)
}
