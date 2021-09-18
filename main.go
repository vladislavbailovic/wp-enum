package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type EnumerationType int

const (
	ENUM_JSON_API EnumerationType = iota
	ENUM_JSON_ROUTE
	ENUM_AUTHOR_ID
)

type apiResponse struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
}

func getJson(apiUrl string) ([]apiResponse, error) {
	client := NewHttpClient(CLIENT_WEB)
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

func enumerateJsonRoute(url string) (map[string]int, error) {
	apiUrl := fmt.Sprintf("%s?rest_route=/wp/v2/users/", url)
	return getJsonUsers(apiUrl)
}

func enumerateAuthorId(url string) (map[string]int, error) {
	result := make(map[string]int)
	client := NewHttpClient(CLIENT_WEB)
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

func Enumerate(kind EnumerationType, url string) (map[string]int, error) {
	if ENUM_JSON_API == kind {
		return enumerateJsonApi(url)
	}
	if ENUM_JSON_ROUTE == kind {
		return enumerateJsonRoute(url)
	}
	if ENUM_AUTHOR_ID == kind {
		return enumerateAuthorId(url)
	}
	return nil, errors.New(fmt.Sprintf("Unknown enumeration type: %d", kind))
}

type ClientType string

const (
	CLIENT_PASSTHROUGH ClientType = "null"
	CLIENT_WEB         ClientType = "web"
)

type Client interface {
	Send(*http.Request) http.Response
}

func NewHttpClient(ctype ...ClientType) Client {
	var clientType ClientType
	if len(ctype) > 0 {
		clientType = ctype[0]
	}

	if CLIENT_WEB == clientType {
		client := &http.Client{
			Timeout: time.Second * 10,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		return WebClient{http: client}
	}
	return PassthroughClient{}
}

type PassthroughClient struct{}

func (nc PassthroughClient) Send(req *http.Request) http.Response {
	return http.Response{StatusCode: -1}
}

type WebClient struct {
	http *http.Client
}

func (wc WebClient) Send(req *http.Request) http.Response {
	response, err := wc.http.Do(req)
	if err != nil {
		return http.Response{StatusCode: -1}
	}
	return *response
}
