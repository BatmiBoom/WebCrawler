package parseurl

import (
	"strings"

	"golang.org/x/net/html"
)

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
