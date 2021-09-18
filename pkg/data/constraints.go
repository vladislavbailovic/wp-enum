package data

type Constraints struct {
	URL   string
	Kind  int
	Start int
	End   int
}

func DefaultConstraints() Constraints {
	return Constraints{
		URL:   "",
		Kind:  0,
		Start: 1,
		End:   10,
	}
}
