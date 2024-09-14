package crawler

import (
	"fmt"
	"net/url"
	"sync"

	handlehtml "github.com/BatmiBoom/web_crawler_go/cmd/handle_html"
	"github.com/BatmiBoom/web_crawler_go/cmd/normalize"
)

type Config struct {
	Pages              map[string]int
	MaxPages           int
	MaxConcurrency     int
	BaseURL            *url.URL
	Mu                 *sync.Mutex
	ConcurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
}

func (cfg *Config) CrawlPage(
	parsedCurrentURL *url.URL,
) (map[string]int, error) {
	fmt.Println("Checking if we suprass max pages")
	if len(cfg.Pages) > cfg.MaxPages {
		return cfg.Pages, nil
	}

	if cfg.BaseURL.Host != parsedCurrentURL.Host {
		return cfg.Pages, fmt.Errorf("ERROR: out of domain")
	}

	normalizeCurrent, err := normalize.NormalizeURL(parsedCurrentURL)
	if err != nil {
		fmt.Println(parsedCurrentURL)
		return cfg.Pages, fmt.Errorf("error: normalizing url %v\n", err)
	}

	_, ok := cfg.Pages[normalizeCurrent]
	if ok {
		fmt.Printf("Page already existed %v\n", normalizeCurrent)
		cfg.Pages[normalizeCurrent]++
		return cfg.Pages, nil
	}

	fmt.Printf("New url %v\n", normalizeCurrent)
	cfg.Pages[normalizeCurrent] = 1

	fmt.Printf("Requesting HTML %v\n", normalizeCurrent)
	rawHTML, err := handlehtml.GetHTML(normalizeCurrent)
	if err != nil {
		return cfg.Pages, fmt.Errorf("ERROR: Getting the HTML %v\n", err)
	}

	fmt.Printf("Parsing links %v\n", normalizeCurrent)
	links, err := handlehtml.GetURLsFromHTML(rawHTML, parsedCurrentURL)
	if err != nil {
		return cfg.Pages, fmt.Errorf("ERROR: Getting the URLS %v\n", err)
	}

	for _, link := range links {
		parsedNewCurrentURL, err := url.Parse(link)
		if err != nil {
			return cfg.Pages, fmt.Errorf("ERROR: Parsing new link %v\n", err)
		}

		fmt.Printf("New Crawl on %v\n", parsedNewCurrentURL)
		cfg.CrawlPage(parsedNewCurrentURL)
	}

	return cfg.Pages, nil
}
