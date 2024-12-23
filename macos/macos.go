package macos

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jcbhmr/go-platformdirs/api"
)

func mustUserHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("os.UserHomeDir() failed: %w", err))
	}
	return homeDir
}

type MacOS interface {
	api.PlatformDirsABC
}

type MacOSImpl struct {
	api.PlatformDirsABCImpl
}

var _ MacOS = (*MacOSImpl)(nil)

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) *MacOSImpl {
	this := &MacOSImpl{*api.NewPlatformDirsABC(appname, appauthor, version, roaming, multipath, opinion, ensureExists)}
	this.This = this
	return this
}

func (p *MacOSImpl) UserDataDir() string {
	return p.This.(MacOS).X__AppendAppNameAndVersion(filepath.Join(mustUserHomeDir(), "Library/Application Support"))
}

func (p *MacOSImpl) SiteDataDir() string {
	var isHomebrew bool
	if exe, err := os.Executable(); err == nil {
		isHomebrew = strings.HasPrefix(exe, "/opt/homebrew")
	} else {
		isHomebrew = false
	}

	var pathList []string
	if isHomebrew {
		pathList = []string{"/opt/homebrew/share"}
	} else {
		pathList = []string{}
	}
	pathList = append(pathList, p.This.(MacOS).X__AppendAppNameAndVersion("/Library/Application Support"))
	if p.Multipath {
		return strings.Join(pathList, string(filepath.ListSeparator))
	} else {
		return pathList[0]
	}
}

func (p *MacOSImpl) SiteDataPath() string {
	return p.This.(MacOS).X__FirstItemAsPathIfMultipath(p.SiteDataDir())
}

func (p *MacOSImpl) UserConfigDir() string {
	return p.This.(MacOS).UserDataDir()
}

func (p *MacOSImpl) SiteConfigDir() string {
	return p.This.(MacOS).SiteDataDir()
}

func (p *MacOSImpl) UserCacheDir() string {
	return p.This.(MacOS).X__AppendAppNameAndVersion(filepath.Join(mustUserHomeDir(), "Library/Caches"))
}

func (p *MacOSImpl) SiteCacheDir() string {
	var isHomebrew bool
	if exe, err := os.Executable(); err == nil {
		isHomebrew = strings.HasPrefix(exe, "/opt/homebrew")
	} else {
		isHomebrew = false
	}

	var pathList []string
	if isHomebrew {
		pathList = []string{"/opt/homebrew/var/cache"}
	} else {
		pathList = []string{}
	}
	pathList = append(pathList, p.This.(MacOS).X__AppendAppNameAndVersion("/Library/Caches"))
	if p.Multipath {
		return strings.Join(pathList, string(filepath.ListSeparator))
	} else {
		return pathList[0]
	}
}

func (p *MacOSImpl) SiteCachePath() string {
	return p.This.(MacOS).X__FirstItemAsPathIfMultipath(p.SiteCacheDir())
}

func (p *MacOSImpl) UserStateDir() string {
	return p.This.(MacOS).UserDataDir()
}

func (p *MacOSImpl) UserLogDir() string {
	return p.This.(MacOS).X__AppendAppNameAndVersion(filepath.Join(mustUserHomeDir(), "Library/Logs"))
}

func (p *MacOSImpl) UserDocumentsDir() string {
	return filepath.Join(mustUserHomeDir(), "Documents")
}

func (p *MacOSImpl) UserDownloadsDir() string {
	return filepath.Join(mustUserHomeDir(), "Downloads")
}

func (p *MacOSImpl) UserPicturesDir() string {
	return filepath.Join(mustUserHomeDir(), "Pictures")
}

func (p *MacOSImpl) UserVideosDir() string {
	return filepath.Join(mustUserHomeDir(), "Movies")
}

func (p *MacOSImpl) UserMusicDir() string {
	return filepath.Join(mustUserHomeDir(), "Music")
}

func (p *MacOSImpl) UserDesktopDir() string {
	return filepath.Join(mustUserHomeDir(), "Desktop")
}

func (p *MacOSImpl) UserRuntimeDir() string {
	return p.This.(MacOS).X__AppendAppNameAndVersion(filepath.Join(mustUserHomeDir(), "Library/Caches/TemporaryItems"))
}

func (p *MacOSImpl) SiteRuntimeDir() string {
	return p.This.(MacOS).UserRuntimeDir()
}
