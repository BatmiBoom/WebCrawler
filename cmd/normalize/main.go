package normalize

import (
	"net/url"
)

func NormalizeURL(uri *url.URL) (string, error) {
	return uri.Scheme + "://" + uri.Host + uri.Path, nil
}
