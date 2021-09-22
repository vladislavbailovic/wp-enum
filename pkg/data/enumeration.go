package data

import wp_http "wp-user-enum/pkg/http"

type EnumerationType int

type Enumerator func(wp_http.Client, Constraints) ([]ApiResponse, error)

const (
	ENUM_JSON_API EnumerationType = iota
	ENUM_JSON_ROUTE
	ENUM_AUTHOR_ID
)
