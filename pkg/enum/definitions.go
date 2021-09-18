package enum

import (
	"errors"
	"fmt"

	wp_http "wp-enum/pkg/http"
)

type EnumerationType int

const (
	TYPE_JSON_API EnumerationType = iota
	TYPE_JSON_ROUTE
	TYPE_AUTHOR_ID
)

type apiResponse struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func Enumerate(kind EnumerationType, url string) (func(wp_http.Client, ...int) (map[string]int, error), error) {
	if TYPE_JSON_API == kind {
		return enumerateJsonApi(url), nil
	}
	if TYPE_JSON_ROUTE == kind {
		return enumerateJsonRoute(url), nil
	}
	if TYPE_AUTHOR_ID == kind {
		return enumerateAuthorId(url), nil
	}
	return nil, errors.New(fmt.Sprintf("Unknown enumeration type: %d", kind))
}
