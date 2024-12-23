package version

import (
	"regexp"
	"strconv"
)

// FIXME: Change this number for each release.

var Version = "4.3.6"

var VersionTuple struct{ A, B, C uint }

func init() {
	match := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)`).FindStringSubmatch(Version)
	if match == nil {
		panic("invalid version")
	}
	a, err := strconv.ParseUint(match[1], 10, 32)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseUint(match[2], 10, 32)
	if err != nil {
		panic(err)
	}
	c, err := strconv.ParseUint(match[3], 10, 32)
	if err != nil {
		panic(err)
	}
	VersionTuple = struct{ A, B, C uint }{uint(a), uint(b), uint(c)}
}
