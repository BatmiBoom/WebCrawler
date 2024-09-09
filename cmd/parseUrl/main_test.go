package parseurl

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "Get URLS from HTML",
			inputURL: "https://blog.boot.dev",
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
			actual, err := GetURLsFromHTML(tc.inputBody, tc.inputURL)
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
