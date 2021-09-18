package enum

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	wp_http "wp-enum/pkg/http"
)

func enumerateAuthorId(url string) func(wp_http.Client, ...int) (map[string]int, error) {
	return func(client wp_http.Client, limit ...int) (map[string]int, error) {
		maxUsers := 100
		if len(limit) > 0 {
			maxUsers = limit[0]
		}
		result := make(map[string]int)
		var overallErr error
		for i := 1; i < maxUsers; i++ {
			author_url := fmt.Sprintf("%s?author=%d", url, i)
			req, _ := http.NewRequest("GET", author_url, nil)
			resp := client.Send(req)
			if resp.StatusCode < 100 {
				overallErr = errors.New("passthrough")
			} else if resp.StatusCode > 300 && resp.StatusCode < 400 {
				location := resp.Header.Get("location")
				rpl := fmt.Sprintf("%sauthor/", url)
				user := strings.Replace(strings.Replace(location, rpl, "", 1), "/", "", -1)
				result[user] = i
			}
		}
		if overallErr != nil && len(result) == 0 {
			return nil, overallErr
		}
		return result, nil
	}
}
