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

	GetAgent() *UserAgent
	SetAgent(ua *UserAgent)

	GetCookies() []*http.Cookie
	HasCookies() bool
	AddCookie(*http.Cookie)
}

func NewHttpClient(ctype ...ClientType) Client {
	var clientType ClientType
	if len(ctype) > 0 {
		clientType = ctype[0]
	}

	if CLIENT_WEB == clientType {
		ua := UserAgent{}
		client := &http.Client{
			Timeout: time.Second * 10,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		return WebClient{http: client, ua: &ua, cookies: &CookieStore{}}
	}
	return PassthroughClient{}
}

type PassthroughClient struct{}

func (nc PassthroughClient) Send(url string) http.Response {
	return http.Response{StatusCode: -1}
}
func (nc PassthroughClient) SetAgent(ua *UserAgent) {}
func (nc PassthroughClient) GetAgent() *UserAgent {
	return &UserAgent{}
}
func (nc PassthroughClient) AddCookie(c *http.Cookie) {}
func (nc PassthroughClient) HasCookies() bool {
	return false
}
func (nc PassthroughClient) GetCookies() []*http.Cookie {
	return []*http.Cookie{}
}

type WebClient struct {
	http    *http.Client
	ua      *UserAgent
	cookies *CookieStore
}

func (wc WebClient) Send(url string) http.Response {
	request, err := http.NewRequest("GET", NormalizeUrl(url), nil)
	if err != nil {
		return http.Response{StatusCode: -1}
	}

	wc.ua.SetHeader(request)
	response, err := wc.http.Do(request)
	if err != nil {
		return http.Response{StatusCode: -1}
	}
	return *response
}
func (wc WebClient) SetAgent(ua *UserAgent) {
	wc.ua.isRandom = ua.isRandom
}
func (wc WebClient) GetAgent() *UserAgent {
	return wc.ua
}
func (wc WebClient) AddCookie(c *http.Cookie) {
	wc.cookies.store = append(wc.cookies.store, c)
}
func (wc WebClient) GetCookies() []*http.Cookie {
	return wc.cookies.store
}
func (wc WebClient) HasCookies() bool {
	return len(wc.cookies.store) != 0
}
