package enum

import (
	"errors"
	"fmt"
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

func Enumerate(kind EnumerationType, url string) (map[string]int, error) {
	if TYPE_JSON_API == kind {
		return enumerateJsonApi(url)
	}
	if TYPE_JSON_ROUTE == kind {
		return enumerateJsonRoute(url)
	}
	if TYPE_AUTHOR_ID == kind {
		return enumerateAuthorId(url)
	}
	return nil, errors.New(fmt.Sprintf("Unknown enumeration type: %d", kind))
}
