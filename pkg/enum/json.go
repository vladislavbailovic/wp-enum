package enum

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	wp_http "wp-enum/pkg/http"
)

func getJson(apiUrl string) ([]apiResponse, error) {
	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	resp := client.Send(req)
	if 200 != resp.StatusCode {
		return nil, errors.New(fmt.Sprintf("non-200 status code contacting JSON API at %s: %d", apiUrl, resp.StatusCode))
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tmp []apiResponse
	err = json.Unmarshal(bodyBytes, &tmp)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}

func getJsonUsers(apiUrl string) (map[string]int, error) {
	result := make(map[string]int)
	json, err := getJson(apiUrl)
	if err != nil {
		return nil, err
	}
	for _, user := range json {
		result[user.Name] = user.Id
	}
	return result, nil
}

func enumerateJsonApi(url string) (map[string]int, error) {
	apiUrl := fmt.Sprintf("%swp-json/wp/v2/users", url)
	return getJsonUsers(apiUrl)
}
