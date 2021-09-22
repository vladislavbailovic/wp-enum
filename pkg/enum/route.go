package enum

import (
	"fmt"
	"wp-user-enum/pkg/data"
	wp_http "wp-user-enum/pkg/http"
)

func enumerateJsonRoute(url string) func(wp_http.Client, data.Constraints) ([]data.ApiResponse, error) {
	apiUrl := fmt.Sprintf("%s?rest_route=/wp/v2/users/", wp_http.NormalizeRootUrl(url))
	return func(client wp_http.Client, limit data.Constraints) ([]data.ApiResponse, error) {
		return getJsonUsers(apiUrl, client)
	}
}
