package data

type Constraints struct {
	URL   string
	Kind  int
	Limit int
}

func DefaultConstraints() Constraints {
	return Constraints{
		URL:   "",
		Kind:  0,
		Limit: 10,
	}
}
