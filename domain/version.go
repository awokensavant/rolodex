package domain

type Version int

const (
	Version4 Version = iota
	Version3
	Version2_1
	UnknownVersion
)

var versionName = map[Version]string{
	Version4:       "4.0.0",
	Version3:       "3.0.0",
	Version2_1:     "2.1.0",
	UnknownVersion: "unknown",
}

func (v Version) String() string {
	switch v {
	case Version4:
		return versionName[Version4]
	case Version3:
		return versionName[Version3]
	case Version2_1:
		return versionName[Version2_1]
	default:
		return versionName[UnknownVersion]
	}
}
