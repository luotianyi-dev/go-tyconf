package tyconf

import (
	"strings"
	_ "time/tzdata"
)

var (
	ModeDevelopment = "development"
	ModeProduction  = "production"

	BuildVersion = "dev"
	BuildTime    = "<unknown>"
	BuildMode    = ModeDevelopment
)

func IsDevelopment() bool {
	return strings.EqualFold(BuildMode, ModeDevelopment)
}
