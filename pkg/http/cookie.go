package http

import (
	"fmt"
	"net/http"
	"strings"
)

type CookieType string

const (
	COOKIE_WP_TEST      CookieType = "wordpress_test_cookie"
	COOKIE_WP_GENERIC   CookieType = "wordpress"
	COOKIE_WP_LOGGED_IN CookieType = "wordpress_logged_in"
	COOKIE_WP_SEC       CookieType = "wordpress_sec"
)

type CookieStore struct {
	store []*http.Cookie
}

type WPCookie struct {
	hash string
}

func (wpc WPCookie) Get(prefix CookieType, rawValue ...string) *http.Cookie {
	var value string
	name := fmt.Sprintf("%s_%s", string(prefix), wpc.hash)
	if len(rawValue) == 1 {
		value = rawValue[0]
	}
	if COOKIE_WP_TEST == prefix {
		name = string(COOKIE_WP_TEST)
		value = "WP Cookie check"
	}
	return &http.Cookie{Name: name, Value: value}
}

func addClientCookies(client Client, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		client.AddCookie(cookie)
	}
}

func AddMockWPCookies(client Client) {
	wpc := WPCookie{strings.Repeat("d3adb3af", 4)}
	mocks := []*http.Cookie{
		wpc.Get(COOKIE_WP_TEST),
		wpc.Get(COOKIE_WP_GENERIC),
		wpc.Get(COOKIE_WP_LOGGED_IN, "mock"),
		wpc.Get(COOKIE_WP_SEC, "mock"),
	}
	addClientCookies(client, mocks)
}
