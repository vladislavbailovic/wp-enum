package data

type Constraints struct {
	URL      string
	Kind     int
	Start    int
	End      int
	RandomUA bool
}

func DefaultConstraints() Constraints {
	return Constraints{
		URL:      "",
		Kind:     0,
		Start:    1,
		End:      10,
		RandomUA: false,
	}
}
