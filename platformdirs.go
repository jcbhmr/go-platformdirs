package platformdirs

import "github.com/jcbhmr/go-platformdirs/api"

type PlatformDirsABC = api.PlatformDirsABC

type AppDirs = PlatformDirs

func NewAppDirs(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) PlatformDirs {
	return New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
}

func UserDataDir(appname *string, appauthor any, version *string, roaming *bool, ensureExists *bool) (string, error) {
	return New(appname, appauthor, version, roaming, nil, nil, ensureExists).UserDataDir()
}

func SiteDataDir(appname *string, appauthor any, version *string, multipath *bool, ensureExists *bool) (string, error) {
	return New(appname, appauthor, version, nil, multipath, nil, ensureExists).SiteDataDir()
}

func UserConfigDir(appname *string, appauthor any, version *string, roaming *bool, ensureExists *bool) (string, error) {
	return New(appname, appauthor, version, roaming, nil, nil, ensureExists).UserConfigDir()
}

func SiteConfigDir(appname *string, appauthor any, version *string, multipath *bool, ensureExists *bool) (string, error) {
	return New(appname, appauthor, version, nil, multipath, nil, ensureExists).SiteConfigDir()
}