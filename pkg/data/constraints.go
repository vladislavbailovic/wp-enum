package data

import "wp-enum/pkg/enum"

type Constraints struct {
	URL   string
	Kind  int
	Limit int
}

func DefaultConstraints() Constraints {
	return Constraints{
		URL:   "",
		Kind:  int(enum.TYPE_JSON_API),
		Limit: 10,
	}
}
