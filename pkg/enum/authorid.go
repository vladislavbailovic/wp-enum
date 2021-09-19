package enum

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"wp-enum/pkg/data"
	wp_http "wp-enum/pkg/http"
)

type authorReq struct {
	url string
	id  int
}

func makeRequests(base string, start, end int) []authorReq {
	urls := make([]authorReq, end-start)

	count := 0
	for idx := start; idx < end; idx++ {
		url := fmt.Sprintf("%s?author=%d", base, idx)
		urls[count] = authorReq{url, idx}
		count++
	}

	return urls
}

func isRedirect(resp http.Response) bool {
	return resp.StatusCode > 300 && resp.StatusCode < 399
}

func getAuthorFromResponse(author_url string, resp http.Response) (string, error) {
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
		if strings.HasPrefix(location, rpl) {
			rawUser = strings.Replace(location, rpl, "", 1)
		}
	}
	return strings.Replace(rawUser, "/", "", -1), nil
}

func getAuthor(author_url string, client wp_http.Client) (string, error) {
	resp := client.Send(author_url)

	user, err := getAuthorFromResponse(author_url, resp)
	if err != nil {
		return "", err
	}

	return user, nil
}

func getAuthorsFromBatch(requestList []authorReq, client wp_http.Client) []data.ApiResponse {
	result := []data.ApiResponse{}
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(requestList))

	for _, author := range requestList {
		go func(author authorReq, client wp_http.Client) {
			defer wg.Done()
			user, err := getAuthor(author.url, client)
			if err != nil {
				return
			}

			if "" != user {
				mu.Lock()
				result = append(result, data.ApiResponse{user, author.id})
				mu.Unlock()
			}
		}(author, client)
	}

	wg.Wait()

	return result
}

func getAllAuthors(requestList []authorReq, client wp_http.Client) ([]data.ApiResponse, error) {
	results := []data.ApiResponse{}

	batch := []authorReq{}
	count := 0
	for _, url := range requestList {
		count++
		if count > 5 {
			for _, author := range getAuthorsFromBatch(batch, client) {
				results = append(results, author)
			}
			count = 0
			batch = []authorReq{}
		}
		batch = append(batch, url)
	}

	if len(batch) > 0 {
		for _, author := range getAuthorsFromBatch(batch, client) {
			results = append(results, author)
		}
	}

	return results, nil
}

func enumerateAuthorId(url string) func(wp_http.Client, data.Constraints) ([]data.ApiResponse, error) {
	url = wp_http.NormalizeRootUrl(url)

	return func(client wp_http.Client, opts data.Constraints) ([]data.ApiResponse, error) {
		requestList := makeRequests(url, opts.Start, opts.End)
		return getAllAuthors(requestList, client)
	}
}
