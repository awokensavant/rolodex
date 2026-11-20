package rolodex

import (
	"fmt"
)

var (
	Version   string
	Revision  string
	BuildTime string
	API       string
)

func GetVersion() string {
	return fmt.Sprintf("go-gedcom v%s commit:%s asof %s satisfying api v%s", Version, Revision, BuildTime, API)
}
