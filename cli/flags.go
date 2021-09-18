package cli

import (
	"flag"
	"wp-enum/pkg/enum"
)

type Options struct {
	Url   string
	Kind  enum.EnumerationType
	Limit int
}

func GetFlags() Options {
	url := flag.String("url", "", "WordPress URL")
	kind := flag.Int("enum", int(enum.TYPE_JSON_API), "Enumeration type")
	limit := flag.Int("limit", 100, "Limit to number of users")
	flag.Parse()

	return Options{*url, enum.EnumerationType(*kind), *limit}
}
