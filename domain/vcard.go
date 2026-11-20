// Package domain contains the application's domain models.
package domain

// Property represents a vCard content-line property. Value is the property's
// RFC 6350 value text, and Parameters holds its content-line parameters.
type Property struct {
	Name       string
	Group      string
	Value      string
	Parameters map[string][]string
}

// VCard represents a vCard 4.0 object as defined by RFC 6350.
// The property slices allow repeated properties; singular RFC properties use
// pointers. Version is the value of the required VERSION property.
type VCard struct {
	Version string

	Source             []Property
	Kind               *Property
	XML                []Property
	FormattedName      []Property
	Name               *Property
	Nickname           []Property
	Photo              []Property
	Birthday           []Property
	Anniversary        []Property
	Gender             *Property
	Address            []Property
	Telephone          []Property
	Email              []Property
	InstantMessaging   []Property
	Language           []Property
	TimeZone           []Property
	Geo                []Property
	Title              []Property
	Role               []Property
	Logo               []Property
	Organization       []Property
	Member             []Property
	Related            []Property
	Categories         []Property
	Note               []Property
	ProductID          *Property
	Revision           *Property
	Sound              []Property
	UID                *Property
	ClientPIDMap       []Property
	URL                []Property
	Key                []Property
	FBURL              []Property
	CalendarAddressURI []Property
	CalendarURI        []Property

	// Extensions retains properties not defined by RFC 6350, including
	// extension properties whose names begin with "X-".
	Extensions []Property
}
