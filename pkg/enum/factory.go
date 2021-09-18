package enum

import (
	"errors"
	"fmt"

	"wp-enum/pkg/data"
)

func Enumerate(kind data.EnumerationType, url string) (data.Enumerator, error) {
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
