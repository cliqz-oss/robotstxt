package robotstxt

import (
	"net/url"
)

// Try to unquote all paths and urls from url encoding.
// On failure, return the path as it is.

func EscapeQuotes(path string) string {
	unquoted, err := url.QueryUnescape(path)
	if err != nil {
		return path
	}
	return unquoted
}
