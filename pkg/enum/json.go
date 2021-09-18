package enum

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"wp-enum/pkg/data"
	wp_http "wp-enum/pkg/http"
)

func getJson(apiUrl string, client wp_http.Client) ([]apiResponse, error) {
	resp := client.Send(apiUrl)
	if 200 != resp.StatusCode {
		return nil, errors.New(fmt.Sprintf("non-200 status code contacting JSON API at %s: %d", apiUrl, resp.StatusCode))
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var tmp []apiResponse
	err := json.Unmarshal(bodyBytes, &tmp)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}

func getJsonUsers(apiUrl string, client wp_http.Client) (map[string]int, error) {
	result := make(map[string]int)
	json, err := getJson(apiUrl, client)
	if err != nil {
		return nil, err
	}
	for _, user := range json {
		result[user.Name] = user.Id
	}
	return result, nil
}

func enumerateJsonApi(url string) func(wp_http.Client, data.Constraints) (map[string]int, error) {
	apiUrl := fmt.Sprintf("%swp-json/wp/v2/users", wp_http.NormalizeRootUrl(url))
	return func(client wp_http.Client, limit data.Constraints) (map[string]int, error) {
		return getJsonUsers(apiUrl, client)
	}
}
