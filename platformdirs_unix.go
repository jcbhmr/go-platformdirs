//go:build unix && !darwin

package platformdirs

import "github.com/jcbhmr/go-platformdirs/unix"

type PlatformDirs = unix.Unix

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) unix.Unix {
	return unix.New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
}