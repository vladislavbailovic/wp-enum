package http

import (
	"net/http"
	"time"
)

type ClientType string

const (
	CLIENT_PASSTHROUGH ClientType = "null"
	CLIENT_WEB         ClientType = "web"
)

type Client interface {
	Send(string) http.Response
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

func (nc PassthroughClient) Send(url string) http.Response {
	return http.Response{StatusCode: -1}
}

type WebClient struct {
	http *http.Client
}

func (wc WebClient) Send(url string) http.Response {
	request, err := http.NewRequest("GET", NormalizeUrl(url), nil)
	if err != nil {
		return http.Response{StatusCode: -1}
	}
	response, err := wc.http.Do(request)
	if err != nil {
		return http.Response{StatusCode: -1}
	}
	return *response
}
