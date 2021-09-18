package enum

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"wp-enum/pkg/data"
	wp_http "wp-enum/pkg/http"
)

func generateUrls(base string, start, end int) []string {
	urls := make([]string, end-start)

	count := 0
	for idx := start; idx < end; idx++ {
		url := fmt.Sprintf("%s?author=%d", base, idx)
		urls[count] = url
		count++
	}

	return urls
}

func isRedirect(resp http.Response) bool {
	return resp.StatusCode > 300 && resp.StatusCode < 399
}

func getAuthorFromRedirect(author_url string, resp http.Response) (string, error) {
	if !isRedirect(resp) {
		return "", errors.New("not a redirect")
	}
	location := resp.Header.Get("location")
	authorParam := strings.LastIndex(location, "/author/")
	var rawUser string
	if authorParam >= 0 {
		rawUser = location[authorParam+8:]
	} else {
		rpl := fmt.Sprintf("%sauthor/", author_url)
		rawUser = strings.Replace(location, rpl, "", 1)
	}
	return strings.Replace(rawUser, "/", "", -1), nil
}

func getUserDataFromUrls(urlList []string, client wp_http.Client) ([]data.ApiResponse, error) {
	result := []data.ApiResponse{}
	var overallErr error

	for idx, author_url := range urlList {
		resp := client.Send(author_url)
		if !isRedirect(resp) {
			continue
		}

		user, err := getAuthorFromRedirect(author_url, resp)
		if err != nil {
			overallErr = err
			continue
		}

		if "" != user {
			result = append(result, data.ApiResponse{user, idx})
		}
	}

	if overallErr != nil && len(result) == 0 {
		return nil, overallErr
	}
	return result, nil
}

func enumerateAuthorId(url string) func(wp_http.Client, data.Constraints) ([]data.ApiResponse, error) {
	url = wp_http.NormalizeRootUrl(url)

	return func(client wp_http.Client, opts data.Constraints) ([]data.ApiResponse, error) {
		urlList := generateUrls(url, opts.Start, opts.End)
		return getUserDataFromUrls(urlList, client)
	}
}
