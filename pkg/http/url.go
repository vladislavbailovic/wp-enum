package http

import (
	"fmt"
	"strings"
)

func isHttp(url string) bool {
	return strings.HasPrefix(url, "http://")
}

func isHttps(url string) bool {
	return strings.HasPrefix(url, "https://")
}

func HasProtocol(url string) bool {
	return isHttp(url) || isHttps(url)
}

func HasRelativeProtocol(url string) bool {
	return strings.HasPrefix(url, "//")
}

func NormalizeUrl(raw string) string {
	if HasProtocol(raw) {
		return raw
	}
	if HasRelativeProtocol(raw) {
		return fmt.Sprintf("http:%s", raw)
	}
	return fmt.Sprintf("http://%s", raw)
}

func NormalizeRootUrl(raw string) string {
	url := NormalizeUrl(raw)
	return Trailingslash(url)
}

func Trailingslash(raw string) string {
	return fmt.Sprintf("%s/", Untrailingslash(raw))
}

func Untrailingslash(raw string) string {
	for strings.HasSuffix(raw, "/") {
		raw = raw[0 : len(raw)-1]
	}
	return raw
}
