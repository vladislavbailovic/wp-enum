package enum

import "fmt"

func enumerateJsonRoute(url string) (map[string]int, error) {
	apiUrl := fmt.Sprintf("%s?rest_route=/wp/v2/users/", url)
	return getJsonUsers(apiUrl)
}
