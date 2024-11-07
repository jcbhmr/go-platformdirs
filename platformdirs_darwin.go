package platformdirs

import "github.com/jcbhmr/go-platformdirs/macos"

type PlatformDirs = macos.MacOS

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) macos.MacOS {
	return android.New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
}