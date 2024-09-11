package handlehtml

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func GetHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("ERROR: There was an error in the request %v", err)
	}

	if resp.StatusCode > 399 {
		return "", fmt.Errorf("ERROR: Status Code %v", resp.StatusCode)
	}

	if !strings.ContainsAny(resp.Header["Content-Type"][0], "text/html") {
		return "", fmt.Errorf("ERROR: Wrong content type %v", resp.Header["Content-Type"][0])	
	}

	rawHTML, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ERROR: There was a problem reading the body %v", err)
	}

	return string(rawHTML), nil
}

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	tokens := html.NewTokenizer(strings.NewReader(htmlBody))

	links := []string{}
	for {
		tt := tokens.Next()

		switch tt {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken:
			tn, has_attr := tokens.TagName()

			if has_attr == false {
				break
			}

			if len(tn) == 1 && tn[0] == 'a' {
				link := getLink(tokens)
				if link == "" {
					break
				}

				links = append(links, link)
			}
		}
	}
}

func getLink(tokenizer *html.Tokenizer) string {
	for {
		attr, attr_values, has_more := tokenizer.TagAttr()
		if string(attr[:]) == "href" {
			return string(attr_values[:])
		}

		if has_more == false {
			return ""
		}

	}
}
