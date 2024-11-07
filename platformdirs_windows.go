package platformdirs

import "github.com/jcbhmr/go-platformdirs/windows"

type PlatformDirs = windows.Windows

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) windows.Windows {
	return windows.New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
}