/*
platformdirs is a library to determine platform-specific system directories. This includes directories where to place cache files, user data, configuration, etc.

The source code and issue tracker are both hosted on [GitHub].

Utilities for determining application-specific dirs.

See https://github.com/jcbhmr/go-platformdirs for details and usage.

[GitHub]: https://github.com/jcbhmr/go-platformdirs
*/

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

// Currently active platform
type PlatformDirs = unix.Unix

// "Windows", "MacOS", "Unix", or "Android".
var platformDirsImplName string

// Currently active platform
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

// Backwards compatibility with appdirs
type AppDirs = PlatformDirs

// Backwards compatibility with appdirs
func NewAppDirs(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) AppDirs {
	return New(appname, appauthor, version, roaming, multipath, opinion, ensureExists)
}

// Params:
//
//   - appname: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appname.
//   - appauthor: Must be nil, string, or false. Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appauthor.
//   - version: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Version.
//   - roaming: Optional. Default is false. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Roaming.
//   - ensureExists: Optional. Default is false. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] EnsureExists.
//
// Returns: data directory tied to the user
func UserDataDir(appname *string, appauthor any, version *string, roaming *bool, ensureExists *bool) string {
	var roaming2 bool
	if roaming != nil {
		roaming2 = *roaming
	} else {
		roaming2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, &roaming2, nil, nil, &ensureExists2).UserDataDir()
}

// Params:
//
//   - appname: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appname.
//   - appauthor: Must be nil, string, or false. Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appauthor.
//   - version: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Version.
//   - multipath: Optional. Default is false. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Multipath.
//   - ensureExists: Optional. Default is false. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] EnsureExists.
//
// Returns: configuration directory shared by users
func SiteDataDir(appname *string, appauthor any, version *string, multipath *bool, ensureExists *bool) string {
	var multipath2 bool
	if multipath != nil {
		multipath2 = *multipath
	} else {
		multipath2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, &multipath2, nil, &ensureExists2).SiteDataDir()
}

// Params:
//
//   - appname: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appname.
//   - appauthor: Must be nil, string, or false. Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appauthor.
//   - version: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Version.
//   - roaming: Optional. Default is false. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Roaming.
//   - ensureExists: Optional. Default is false. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] EnsureExists.
//
// Returns: config directory tied to the user
func UserConfigDir(appname *string, appauthor any, version *string, roaming *bool, ensureExists *bool) string {
	var roaming2 bool
	if roaming != nil {
		roaming2 = *roaming
	} else {
		roaming2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, &roaming2, nil, nil, &ensureExists2).UserConfigDir()
}

// Params:
//
//   - appname: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appname.
//   - appauthor: Must be nil, string, or false. Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appauthor.
//   - version: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Version.
//   - multipath: Optional. Default is false. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Multipath.
//   - ensureExists: Optional. Default is false. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] EnsureExists.
//
// Returns: config directory shared by users
func SiteConfigDir(appname *string, appauthor any, version *string, multipath *bool, ensureExists *bool) string {
	var multipath2 bool
	if multipath != nil {
		multipath2 = *multipath
	} else {
		multipath2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, &multipath2, nil, &ensureExists2).SiteConfigDir()
}

// Params:
//
//   - appname: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appname.
//   - appauthor: Must be nil, string, or false. Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Appauthor.
//   - version: Optional. Default is nil. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Version.
//   - opinion: Optional. Default is true. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] Opinion.
//   - ensureExists: Optional. Default is false. See [github.com/jcbhmr/go-platformdirs/api.PlatformDirsABCImpl] EnsureExists.
//
// Returns: cache directory tied to the user
func UserCacheDir(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).UserCacheDir()
}

func SiteCacheDir(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).SiteCacheDir()
}

func UserStateDir(appname *string, appauthor any, version *string, roaming *bool, ensureExists *bool) string {
	var roaming2 bool
	if roaming != nil {
		roaming2 = *roaming
	} else {
		roaming2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, &roaming2, nil, nil, &ensureExists2).UserStateDir()
}

func UserLogDir(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).UserLogDir()
}

func UserDocumentsDir() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserDocumentsDir()
}

func UserDownloadsDir() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserDownloadsDir()
}

func UserPicturesDir() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserPicturesDir()
}

func UserVideosDir() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserVideosDir()
}

func UserMusicDir() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserMusicDir()
}

func UserDesktopDir() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserDesktopDir()
}

func UserRuntimeDir(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).UserRuntimeDir()
}

func SiteRuntimeDir(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).SiteRuntimeDir()
}

func UserDataPath(appname *string, appauthor any, version *string, roaming *bool, ensureExists *bool) string {
	var roaming2 bool
	if roaming != nil {
		roaming2 = *roaming
	} else {
		roaming2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, &roaming2, nil, nil, &ensureExists2).UserDataPath()
}

func SiteDataPath(appname *string, appauthor any, version *string, multipath *bool, ensureExists *bool) string {
	var multipath2 bool
	if multipath != nil {
		multipath2 = *multipath
	} else {
		multipath2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, &multipath2, nil, &ensureExists2).SiteDataPath()
}

func UserConfigPath(appname *string, appauthor any, version *string, roaming *bool, ensureExists *bool) string {
	var roaming2 bool
	if roaming != nil {
		roaming2 = *roaming
	} else {
		roaming2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, &roaming2, nil, nil, &ensureExists2).UserConfigPath()
}

func SiteConfigPath(appname *string, appauthor any, version *string, multipath *bool, ensureExists *bool) string {
	var multipath2 bool
	if multipath != nil {
		multipath2 = *multipath
	} else {
		multipath2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, &multipath2, nil, &ensureExists2).SiteConfigPath()
}

func UserCachePath(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).UserCachePath()
}

func SiteCachePath(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).SiteCachePath()
}

func UserStatePath(appname *string, appauthor any, version *string, roaming *bool, ensureExists *bool) string {
	var roaming2 bool
	if roaming != nil {
		roaming2 = *roaming
	} else {
		roaming2 = false
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, &roaming2, nil, nil, &ensureExists2).UserStatePath()
}

func UserLogPath(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).UserLogPath()
}

func UserDocumentsPath() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserDocumentsPath()
}

func UserDownloadsPath() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserDownloadsPath()
}

func UserPicturesPath() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserPicturesPath()
}

func UserVideosPath() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserVideosPath()
}

func UserMusicPath() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserMusicPath()
}

func UserDesktopPath() string {
	return New(nil, nil, nil, nil, nil, nil, nil).UserDesktopPath()
}

func UserRuntimePath(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).UserRuntimePath()
}

func SiteRuntimePath(appname *string, appauthor any, version *string, opinion *bool, ensureExists *bool) string {
	var opinion2 bool
	if opinion != nil {
		opinion2 = *opinion
	} else {
		opinion2 = true
	}
	var ensureExists2 bool
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	} else {
		ensureExists2 = false
	}
	return New(appname, appauthor, version, nil, nil, &opinion2, &ensureExists2).SiteRuntimePath()
}

var Version = version.Version

var VersionInfo = version.VersionTuple
