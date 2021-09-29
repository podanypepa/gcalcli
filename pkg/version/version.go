// Package version for injectiong info about git version into binary file
package version

import (
	//nolint
	_ "embed"
)

//go:embed buildinfo.txt
var buildInfo []byte

// GetBuildInfo return info about build
func GetBuildInfo() string {
	return string(buildInfo)
}
