// Base API.

package api

import (
	"iter"
	"os"
	"path/filepath"
)

type PlatformDirsABC interface {
	X__AppendAppNameAndVersion(base ...string) string
	X__OptionallyCreateDirectory(path string)
	X__FirstItemAsPathIfMultipath(directory string) string
	UserDataDir() string
	SiteDataDir() string
	UserConfigDir() string
	SiteConfigDir() string
	UserCacheDir() string
	SiteCacheDir() string
	UserStateDir() string
	UserLogDir() string
	UserDocumentsDir() string
	UserDownloadsDir() string
	UserPicturesDir() string
	UserVideosDir() string
	UserMusicDir() string
	UserDesktopDir() string
	UserRuntimeDir() string
	SiteRuntimeDir() string
	UserDataPath() string
	SiteDataPath() string
	UserConfigPath() string
	SiteConfigPath() string
	UserCachePath() string
	SiteCachePath() string
	UserStatePath() string
	UserLogPath() string
	UserDocumentsPath() string
	UserDownloadsPath() string
	UserPicturesPath() string
	UserVideosPath() string
	UserMusicPath() string
	UserDesktopPath() string
	UserRuntimePath() string
	SiteRuntimePath() string
	IterConfigDirs() iter.Seq[string]
	IterDataDirs() iter.Seq[string]
	IterCacheDirs() iter.Seq[string]
	IterRuntimeDirs() iter.Seq[string]
	IterConfigPaths() iter.Seq[string]
	IterDataPaths() iter.Seq[string]
	IterCachePaths() iter.Seq[string]
	IterRuntimePaths() iter.Seq[string]
}

func NewPlatformDirsABC(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) *PlatformDirsABCImpl {
	switch v := appauthor.(type) {
	case nil:
		break
	case string:
		break
	case bool:
		if v {
			panic("invalid appauthor. must be nil, string, or false")
		}
	default:
		panic("invalid appauthor. must be nil, string, or false")
	}
	var roaming2 bool
	if roaming == nil {
		roaming2 = false
	} else {
		roaming2 = *roaming
	}
	var multipath2 bool
	if multipath == nil {
		multipath2 = false
	} else {
		multipath2 = *multipath
	}
	var opinion2 bool
	if opinion == nil {
		opinion2 = true
	} else {
		opinion2 = *opinion
	}
	var ensureExists2 bool
	if ensureExists == nil {
		ensureExists2 = false
	} else {
		ensureExists2 = *ensureExists
	}

	this := &PlatformDirsABCImpl{
		Appname:      appname,
		Appauthor:    appauthor,
		Version:      version,
		Roaming:      roaming2,
		Multipath:    multipath2,
		Opinion:      opinion2,
		EnsureExists: ensureExists2,
	}
	this.This = this
	return this
}

type PlatformDirsABCImpl struct {
	This         PlatformDirsABC
	Appname      *string
	Appauthor    any
	Version      *string
	Roaming      bool
	Multipath    bool
	Opinion      bool
	EnsureExists bool
}

var _ PlatformDirsABC = (*PlatformDirsABCImpl)(nil)

func (p *PlatformDirsABCImpl) X__AppendAppNameAndVersion(base ...string) string {
	params := base[1:]
	if p.Appname != nil {
		params = append(params, *p.Appname)
		if p.Version != nil {
			params = append(params, *p.Version)
		}
	}
	path := filepath.Join(append([]string{base[0]}, params...)...)
	p.This.X__OptionallyCreateDirectory(path)
	return path
}

func (p *PlatformDirsABCImpl) X__OptionallyCreateDirectory(path string) {
	if p.EnsureExists {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			// TODO: Properly Go-ly bubble error?
			panic(err)
		}
	}
}

func (p *PlatformDirsABCImpl) X__FirstItemAsPathIfMultipath(directory string) string {
	if p.Multipath {
		directory = filepath.SplitList(directory)[0]
	}
	return directory
}

func (p *PlatformDirsABCImpl) UserDataDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) SiteDataDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserConfigDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) SiteConfigDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserCacheDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) SiteCacheDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserStateDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserLogDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserDocumentsDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserDownloadsDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserPicturesDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserVideosDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserMusicDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserDesktopDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserRuntimeDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) SiteRuntimeDir() string {
	panic("abstract")
}

func (p *PlatformDirsABCImpl) UserDataPath() string {
	return p.This.UserDataDir()
}

func (p *PlatformDirsABCImpl) SiteDataPath() string {
	return p.This.SiteDataDir()
}

func (p *PlatformDirsABCImpl) UserConfigPath() string {
	return p.This.UserConfigDir()
}

func (p *PlatformDirsABCImpl) SiteConfigPath() string {
	return p.This.SiteConfigDir()
}

func (p *PlatformDirsABCImpl) UserCachePath() string {
	return p.This.UserCacheDir()
}

func (p *PlatformDirsABCImpl) SiteCachePath() string {
	return p.This.SiteCacheDir()
}

func (p *PlatformDirsABCImpl) UserStatePath() string {
	return p.This.UserStateDir()
}

func (p *PlatformDirsABCImpl) UserLogPath() string {
	return p.This.UserLogDir()
}

func (p *PlatformDirsABCImpl) UserDocumentsPath() string {
	return p.This.UserDocumentsDir()
}

func (p *PlatformDirsABCImpl) UserDownloadsPath() string {
	return p.This.UserDownloadsDir()
}

func (p *PlatformDirsABCImpl) UserPicturesPath() string {
	return p.This.UserPicturesDir()
}

func (p *PlatformDirsABCImpl) UserVideosPath() string {
	return p.This.UserVideosDir()
}

func (p *PlatformDirsABCImpl) UserMusicPath() string {
	return p.This.UserMusicDir()
}

func (p *PlatformDirsABCImpl) UserDesktopPath() string {
	return p.This.UserDesktopDir()
}

func (p *PlatformDirsABCImpl) UserRuntimePath() string {
	return p.This.UserRuntimeDir()
}

func (p *PlatformDirsABCImpl) SiteRuntimePath() string {
	return p.This.SiteRuntimeDir()
}

func (p *PlatformDirsABCImpl) IterConfigDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.UserConfigDir()) {
			return
		}
		if !yield(p.This.SiteConfigDir()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterDataDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.UserDataDir()) {
			return
		}
		if !yield(p.This.SiteDataDir()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterCacheDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.UserCacheDir()) {
			return
		}
		if !yield(p.This.SiteCacheDir()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterRuntimeDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.UserRuntimeDir()) {
			return
		}
		if !yield(p.This.SiteRuntimeDir()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterConfigPaths() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.UserConfigPath()) {
			return
		}
		if !yield(p.This.SiteConfigPath()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterDataPaths() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.UserDataPath()) {
			return
		}
		if !yield(p.This.SiteDataPath()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterCachePaths() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.UserCachePath()) {
			return
		}
		if !yield(p.This.SiteCachePath()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterRuntimePaths() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.UserRuntimePath()) {
			return
		}
		if !yield(p.This.SiteRuntimePath()) {
			return
		}
	}
}

// For compatibility with the "Unix" exported type. This is a hack.
func (p *PlatformDirsABCImpl) X__SiteDataDirs() []string {
	panic("logic error")
}
func (p *PlatformDirsABCImpl) X__SiteConfigDirs() []string {
	panic("logic error")
}
