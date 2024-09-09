package normalize

import (
	"fmt"
	"net/url"
)

func normalizeURL(uri string) (string, error) {
	parsed_url, err := url.Parse(uri)
	if err != nil {
		return "", fmt.Errorf("Error parsing URL %v", err)
	}

	return parsed_url.Host + parsed_url.Path, nil
}
