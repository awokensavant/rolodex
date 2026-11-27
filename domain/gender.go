package domain

type Gender int

const (
	Male Gender = iota
	Female
	Transgender
	Unknown
)

var genderName = map[Gender]string{
	Male:        "male",
	Female:      "female",
	Transgender: "transgender",
	Unknown:     "unknown",
}

func (g Gender) String() string {
	switch g {
	case Male:
		return genderName[Male]
	case Female:
		return genderName[Female]
	case Transgender:
		return genderName[Transgender]
	default:
		return genderName[Unknown]
	}
}
