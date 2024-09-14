package crawler

import (
	"fmt"
	"net/url"

	handlehtml "github.com/BatmiBoom/web_crawler_go/cmd/handle_html"
	"github.com/BatmiBoom/web_crawler_go/cmd/normalize"
)

func CrawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (map[string]int, error) {
	normalizeBase, err := normalize.NormalizeURL(rawBaseURL)
	if err != nil {
		fmt.Println(rawBaseURL)
		return pages, fmt.Errorf("error: normalizing url %v\n", err)
	}

	normalizeCurrent, err := normalize.NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(rawCurrentURL)
		return pages, fmt.Errorf("error: normalizing url %v\n", err)
	}

	// fmt.Printf("Normalizing URLs Base %s Curr %s \n", normalizeBase, normalizeCurrent)

	parsedBase, err := url.Parse(normalizeBase)
	if err != nil {
		return pages, fmt.Errorf("error: parsing the url %v\n", err)
	}

	parsedCurrent, err := url.Parse(normalizeCurrent)
	if err != nil {
		return pages, fmt.Errorf("error: parsing the url %v\n", err)
	}

	if parsedBase.Host == "" || parsedCurrent.Host == "" {
		return pages, fmt.Errorf("ERROR: no host")
	}

	if parsedBase.Host != parsedCurrent.Host {
		return pages, fmt.Errorf("ERROR: out of domain")
	}

	_, ok := pages[normalizeCurrent]
	if ok {
		// fmt.Printf("Alredy exist in the map %s\n", normalizeCurrent)
		pages[normalizeCurrent]++
		return pages, nil
	}

	// fmt.Printf("Adding entry in the map %s\n", normalizeCurrent)
	pages[normalizeCurrent] = 1

	rawHTML, err := handlehtml.GetHTML(rawCurrentURL)
	if err != nil {
		return pages, fmt.Errorf("ERROR: Getting the HTML %v\n", err)
	}

	links, err := handlehtml.GetURLsFromHTML(rawHTML, parsedCurrent)
	if err != nil {
		return pages, fmt.Errorf("ERROR: Getting the URLS %v\n", err)
	}
	// fmt.Printf("Links %v\n", links)

	for _, link := range links {
		CrawlPage(rawBaseURL, link, pages)
	}

	return pages, nil
}
