package domain

import "regexp"

const EMAIL_REGEX = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

type Email struct {
	Address string
}

func (e Email) IsValid() bool {
	re := regexp.MustCompile(EMAIL_REGEX)
	return re.MatchString(e.Address)
}
