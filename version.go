package rolodex

import (
	"fmt"
)

var (
	Name      string
	Version   string
	Revision  string
	BuildTime string
	API       string
)

func GetVersion() string {
	return fmt.Sprintf("%s v%s commit:%s asof %s satisfying api v%s", Name, Version, Revision, BuildTime, API)
}
