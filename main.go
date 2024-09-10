package main

import (
	"fmt"
	"os"
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
}
