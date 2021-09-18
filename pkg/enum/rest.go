package enum

import (
	"fmt"
	"net/http"
	"strings"

	wp_http "wp-enum/pkg/http"
)

func enumerateAuthorId(url string) func(wp_http.Client, ...int) (map[string]int, error) {
	return func(client wp_http.Client, limit ...int) (map[string]int, error) {
		result := make(map[string]int)
		for i := 1; i < 10; i++ {
			author_url := fmt.Sprintf("%s?author=%d", url, i)
			req, err := http.NewRequest("GET", author_url, nil)
			if err != nil {
				continue
			}
			resp := client.Send(req)
			if resp.StatusCode > 300 && resp.StatusCode < 400 {
				location := resp.Header.Get("location")
				rpl := fmt.Sprintf("%sauthor/", url)
				user := strings.Replace(strings.Replace(location, rpl, "", 1), "/", "", -1)
				result[user] = i
			}
		}
		return result, nil
	}
}
