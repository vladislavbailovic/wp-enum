package enum

import (
	"fmt"
	"wp-enum/pkg/data"
	wp_http "wp-enum/pkg/http"
)

func enumerateJsonRoute(url string) func(wp_http.Client, data.Constraints) (map[string]int, error) {
	apiUrl := fmt.Sprintf("%s?rest_route=/wp/v2/users/", wp_http.NormalizeRootUrl(url))
	return func(client wp_http.Client, limit data.Constraints) (map[string]int, error) {
		return getJsonUsers(apiUrl, client)
	}
}
