package robotstxt

import (
	"regexp"
	"strings"
	"github.com/ryanuber/go-glob"
)

const (
	REGEXP_MATCHER MatcherType = iota
	GLOB_MATCHER
)

type MatcherType int

type MatcherConstructor func(string) (Matcher, error)

type Matcher interface {
	MatchString(string) bool
	String() string
}

type RegexpMatcher struct {
	re *regexp.Regexp
}

func NewRegexpMatcher(pattern string) (Matcher, error) {
	// Escape string before compile.
	pattern = regexp.QuoteMeta(pattern)
	pattern = strings.Replace(pattern, `\*`, `.*`, -1)
	pattern = strings.Replace(pattern, `\$`, `$`, -1)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &RegexpMatcher{re}, nil
}

func (matcher *RegexpMatcher) MatchString(s string) bool {
	return matcher.re.MatchString(s)
}

func (matcher *RegexpMatcher) String() string {
	return matcher.re.String()
}

type GlobMatcher struct {
	pattern string
	originalPattern string
}

func NewGlobMatcher(pattern string) (Matcher, error) {
	originalPattern := pattern

	// Glob checks for full string match, so just remove trailing '$'
	if strings.HasSuffix(pattern, "$") {
		pattern = strings.TrimRight(pattern, "$")

	// If no explicit end of line, match for any suffix
	} else if !strings.HasSuffix(pattern, "*") {
		pattern = pattern + "*"
	}

	return &GlobMatcher{pattern, originalPattern}, nil
}

func (matcher *GlobMatcher) MatchString(s string) bool {
	return glob.Glob(matcher.pattern, s)
}

func (matcher *GlobMatcher) String() string {
	// Return regex-style pattern. Used only for tests.
	pattern := matcher.originalPattern
	pattern = regexp.QuoteMeta(pattern)
	pattern = strings.Replace(pattern, `\*`, `.*`, -1)
	pattern = strings.Replace(pattern, `\$`, `$`, -1)
	return pattern
}
