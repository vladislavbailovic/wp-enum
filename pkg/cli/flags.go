package cli

import (
	"flag"
	"wp-enum/pkg/data"
)

func GetFlags() data.Constraints {
	defaults := data.DefaultConstraints()

	url := flag.String("url", defaults.URL, "WordPress URL")
	kind := flag.Int("enum", defaults.Kind, "Enumeration type")
	limit := flag.Int("limit", defaults.Limit, "Limit to number of users")
	flag.Parse()

	defaults.URL = *url
	defaults.Kind = *kind
	defaults.Limit = *limit

	return defaults
}
