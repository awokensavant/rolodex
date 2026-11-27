package domain

import "regexp"

const TELEPHONE_REGEX = `^\+?[0-9\s\-\(\)]+$`

type Telephone struct {
	Number string
	Type   TelephoneType
}

func (t Telephone) IsValid() bool {
	re := regexp.MustCompile(TELEPHONE_REGEX)
	return re.MatchString(t.Number)
}
