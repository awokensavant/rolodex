package domain

import "time"

type Card struct {
	Version     Version
	Source      string
	Kind        string
	Name        Name
	Nickname    string
	Birthday    time.Time
	Anniversary time.Time
	Gender      Gender
	Address     Address
	Telephone   []Telephone
	Email       []Email
}
