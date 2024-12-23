package platformdirs

import (
	"os"
	"runtime"

	"github.com/jcbhmr/go-platformdirs/android"
	"github.com/jcbhmr/go-platformdirs/macos"
	"github.com/jcbhmr/go-platformdirs/unix"
	"github.com/jcbhmr/go-platformdirs/version"
	"github.com/jcbhmr/go-platformdirs/windows"
)

// Typed like "Unix" but is actually a runtime-determined implementation: "Windows", "MacOS", "Unix", or "Android".
type PlatformDirs = unix.Unix

var platformDirsImplName string

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) PlatformDirs {
	switch platformDirsImplName {
	case "Windows":
		return windows.New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
	case "MacOS":
		return macos.New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
	case "Unix":
		return unix.New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
	case "Android":
		return android.New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
	default:
		panic("unreachable")
	}
}

func init() {
	var result string
	if runtime.GOOS == "windows" {
		result = "Windows"
	} else if runtime.GOOS == "darwin" {
		result = "MacOS"
	} else {
		result = "Unix"
	}

	setPlatformDirClass := func() string {
		if os.Getenv("ANDROID_DATA") == "/data" && os.Getenv("ANDROID_ROOT") == "/system" {
			if os.Getenv("SHELL") != "" || os.Getenv("PREFIX") != "" {
				return result
			}

			if _, ok := android.X__AndroidFolder(); ok {
				return "Android"
			}
		}

		return result
	}

	platformDirsImplName = setPlatformDirClass()
}

type AppDirs = PlatformDirs

func NewAppDirs(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) AppDirs {
	return New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
}

func UserDataDir(appname *string, appauthor any, version *string, roaming *bool, ensureExists *bool) string {
	return New(appname, appauthor, version, roaming, nil, nil, ensureExists).UserDataDir()
}

func SiteDataDir(appname *string, appauthor any, version *string, multipath *bool, ensureExists *bool) string {
	return New(appname, appauthor, version, nil, multipath, nil, ensureExists).SiteDataDir()
}

var Version = version.Version
var VersionInfo = version.VersionTuple
