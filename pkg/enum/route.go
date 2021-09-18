package enum

import (
	"fmt"
	wp_http "wp-enum/pkg/http"
)

func enumerateJsonRoute(url string) func(wp_http.Client, ...int) (map[string]int, error) {
	apiUrl := fmt.Sprintf("%s?rest_route=/wp/v2/users/", url)
	return func(client wp_http.Client, limit ...int) (map[string]int, error) {
		return getJsonUsers(apiUrl, client)
	}
}
