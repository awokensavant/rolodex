package domain

type TelephoneType int

const (
	Mobile TelephoneType = iota
	Home
	Work
	Other
)

var telephoneTypeName = map[TelephoneType]string{
	Mobile: "mobile",
	Home:   "home",
	Work:   "work",
	Other:  "other",
}

func (t TelephoneType) String() string {
	switch t {
	case Mobile:
		return telephoneTypeName[Mobile]
	case Home:
		return telephoneTypeName[Home]
	case Work:
		return telephoneTypeName[Work]
	default:
		return telephoneTypeName[Other]
	}
}
