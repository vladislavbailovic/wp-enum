package enum

import (
	"errors"
	"fmt"

	"wp-enum/pkg/data"
	wp_http "wp-enum/pkg/http"
)

type apiResponse struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func Enumerate(kind data.EnumerationType, url string) (func(wp_http.Client, data.Constraints) (map[string]int, error), error) {
	if data.ENUM_JSON_API == kind {
		return enumerateJsonApi(url), nil
	}
	if data.ENUM_JSON_ROUTE == kind {
		return enumerateJsonRoute(url), nil
	}
	if data.ENUM_AUTHOR_ID == kind {
		return enumerateAuthorId(url), nil
	}
	return nil, errors.New(fmt.Sprintf("Unknown enumeration type: %d", kind))
}
