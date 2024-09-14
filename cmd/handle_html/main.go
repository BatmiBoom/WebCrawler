package handlehtml

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func GetURLsFromHTML(htmlBody string, parsedUrl *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var urls []string
	var traverseNodes func(*html.Node)
	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key == "href" {
					href, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", anchor.Val, err)
						continue
					}

					resolvedURL := parsedUrl.ResolveReference(href)
					urls = append(urls, resolvedURL.String())
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}
	}
	traverseNodes(doc)

	return urls, nil
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
