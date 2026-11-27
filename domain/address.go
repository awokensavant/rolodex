package domain

type Address struct {
	POBox      string
	Extended   string
	Street     string
	Locality   string
	Region     string
	PostalCode string
	Country    string
}

// String returns the address as a single string, with components separated by semicolons.
func (a Address) String() string {
	return a.POBox + ";" + a.Extended + ";" + a.Street + ";" + a.Locality + ";" + a.Region + ";" + a.PostalCode + ";" + a.Country
}
