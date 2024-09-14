package handlehtml

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		inputUrl  string
		expected  []string
	}{
		{
			name:     "Get URLS from HTML",
			inputUrl: "https://blog.boot.dev/",
			inputBody: `
				<html>
						<body>
								<a target="_blank" href="https://blog.boot.dev/"><span>Go to Boot.dev, you React Andy</span></a>
								<a target="_blank" href="https://other.com/"><span>Go to Boot.dev, you React Andy</span></a>
						</body>
				</html>`,
			expected: []string{"https://blog.boot.dev/", "https://other.com/"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parsedUrl, _ := url.Parse(tc.inputUrl)
			actual, err := GetURLsFromHTML(tc.inputBody, parsedUrl)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
