package cli

import (
	"flag"
	"wp-user-enum/pkg/data"
)

func GetFlags() data.Constraints {
	defaults := data.DefaultConstraints()

	url := flag.String("url", defaults.URL, "WordPress URL")
	kind := flag.Int("enum", defaults.Kind, "Enumeration type")
	start := flag.Int("start", defaults.Start, "Start enumeration at this user ID")
	end := flag.Int("end", defaults.End, "End enumeration with this user ID")
	random := flag.Bool("ua", defaults.RandomUA, "Randomize User-Agent")
	mock := flag.Bool("cookies", defaults.MockCookies, "Send mock WP cookies. Helps with some modsec rulesets (comodo WAF)")
	pretty := flag.Bool("pretty", defaults.Pretty, "Pretty-print the results")

	waf := flag.Bool("waf", defaults.MockCookies && defaults.RandomUA, "Attempt to work around a WAF (randomizes UA and sends mock WP cookies)")

	flag.Parse()

	defaults.URL = *url
	defaults.Kind = *kind
	defaults.Start = *start
	defaults.End = *end
	defaults.RandomUA = *random
	defaults.MockCookies = *mock
	defaults.Pretty = *pretty

	if *waf {
		defaults.MockCookies = true
		defaults.RandomUA = true
	}

	return defaults
}
