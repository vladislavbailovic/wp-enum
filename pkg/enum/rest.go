package enum

import (
	"errors"
	"fmt"
	"strings"

	"wp-enum/pkg/data"
	wp_http "wp-enum/pkg/http"
)

func enumerateAuthorId(url string) func(wp_http.Client, data.Constraints) (map[string]int, error) {
	url = wp_http.NormalizeRootUrl(url)
	return func(client wp_http.Client, opts data.Constraints) (map[string]int, error) {
		result := make(map[string]int)
		var overallErr error
		for i := 1; i < opts.Limit; i++ {
			author_url := fmt.Sprintf("%s?author=%d", url, i)
			resp := client.Send(author_url)
			if resp.StatusCode < 100 {
				overallErr = errors.New("passthrough")
			} else if resp.StatusCode > 300 && resp.StatusCode < 400 {
				location := resp.Header.Get("location")
				authorParam := strings.LastIndex(location, "/author/")
				var rawUser string
				if authorParam >= 0 {
					rawUser = location[authorParam+8:]
				} else {
					rpl := fmt.Sprintf("%sauthor/", url)
					rawUser = strings.Replace(location, rpl, "", 1)
				}
				user := strings.Replace(rawUser, "/", "", -1)
				result[user] = i
			}
		}
		if overallErr != nil && len(result) == 0 {
			return nil, overallErr
		}
		return result, nil
	}
}
