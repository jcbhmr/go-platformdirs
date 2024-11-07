package platformdirs

import "github.com/jcbhmr/go-platformdirs/android"

type PlatformDirs = android.Android

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) android.Android {
	return android.New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
}