// Base API.

package api

import (
	"iter"
	"os"
	"path/filepath"
)

type PlatformDirsABC interface {
	Appname() (string, bool)
	SetAppname(v *string)
	Appauthor() any
	SetAppauthor(v any)
	Version() (string, bool)
	SetVersion(v *string)
	Roaming() bool
	SetRoaming(v bool)
	Multipath() bool
	SetMultipath(v bool)
	Opinion() bool
	SetOpinion(v bool)
	EnsureExists() bool
	SetEnsureExists(v bool)
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
		appname:      appname,
		appauthor:    appauthor,
		version:      version,
		roaming:      roaming2,
		multipath:    multipath2,
		opinion:      opinion2,
		ensureExists: ensureExists2,
	}
	this.X__This = this
	return this
}

type PlatformDirsABCImpl struct {
	X__This      PlatformDirsABC
	appname      *string
	appauthor    any
	version      *string
	roaming      bool
	multipath    bool
	opinion      bool
	ensureExists bool
}

var _ PlatformDirsABC = (*PlatformDirsABCImpl)(nil)

func (p *PlatformDirsABCImpl) X__AppendAppNameAndVersion(base ...string) string {
	params := base[1:]
	if appname, ok := p.X__This.Appname(); ok {
		params = append(params, appname)
		if version, ok := p.X__This.Version(); ok {
			params = append(params, version)
		}
	}
	path := filepath.Join(append([]string{base[0]}, params...)...)
	p.X__This.X__OptionallyCreateDirectory(path)
	return path
}

func (p *PlatformDirsABCImpl) X__OptionallyCreateDirectory(path string) {
	if p.X__This.EnsureExists() {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			// TODO: Properly Go-ly bubble error?
			panic(err)
		}
	}
}

func (p *PlatformDirsABCImpl) X__FirstItemAsPathIfMultipath(directory string) string {
	if p.X__This.Multipath() {
		directory = filepath.SplitList(directory)[0]
	}
	return directory
}

func (p *PlatformDirsABCImpl) Appname() (string, bool) {
	if p.appname == nil {
		return "", false
	}
	return *p.appname, true
}

func (p *PlatformDirsABCImpl) SetAppname(v *string) {
	p.appname = v
}

func (p *PlatformDirsABCImpl) Appauthor() any {
	return p.appauthor
}

func (p *PlatformDirsABCImpl) SetAppauthor(v any) {
	switch v := v.(type) {
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
	p.appauthor = v
}

func (p *PlatformDirsABCImpl) Version() (string, bool) {
	if p.version == nil {
		return "", false
	}
	return *p.version, true
}

func (p *PlatformDirsABCImpl) SetVersion(v *string) {
	p.version = v
}

func (p *PlatformDirsABCImpl) Roaming() bool {
	return p.roaming
}

func (p *PlatformDirsABCImpl) SetRoaming(v bool) {
	p.roaming = v
}

func (p *PlatformDirsABCImpl) Multipath() bool {
	return p.multipath
}

func (p *PlatformDirsABCImpl) SetMultipath(v bool) {
	p.multipath = v
}

func (p *PlatformDirsABCImpl) Opinion() bool {
	return p.opinion
}

func (p *PlatformDirsABCImpl) SetOpinion(v bool) {
	p.opinion = v
}

func (p *PlatformDirsABCImpl) EnsureExists() bool {
	return p.ensureExists
}

func (p *PlatformDirsABCImpl) SetEnsureExists(v bool) {
	p.ensureExists = v
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
	return p.X__This.UserDataDir()
}

func (p *PlatformDirsABCImpl) SiteDataPath() string {
	return p.X__This.SiteDataDir()
}

func (p *PlatformDirsABCImpl) UserConfigPath() string {
	return p.X__This.UserConfigDir()
}

func (p *PlatformDirsABCImpl) SiteConfigPath() string {
	return p.X__This.SiteConfigDir()
}

func (p *PlatformDirsABCImpl) UserCachePath() string {
	return p.X__This.UserCacheDir()
}

func (p *PlatformDirsABCImpl) SiteCachePath() string {
	return p.X__This.SiteCacheDir()
}

func (p *PlatformDirsABCImpl) UserStatePath() string {
	return p.X__This.UserStateDir()
}

func (p *PlatformDirsABCImpl) UserLogPath() string {
	return p.X__This.UserLogDir()
}

func (p *PlatformDirsABCImpl) UserDocumentsPath() string {
	return p.X__This.UserDocumentsDir()
}

func (p *PlatformDirsABCImpl) UserDownloadsPath() string {
	return p.X__This.UserDownloadsDir()
}

func (p *PlatformDirsABCImpl) UserPicturesPath() string {
	return p.X__This.UserPicturesDir()
}

func (p *PlatformDirsABCImpl) UserVideosPath() string {
	return p.X__This.UserVideosDir()
}

func (p *PlatformDirsABCImpl) UserMusicPath() string {
	return p.X__This.UserMusicDir()
}

func (p *PlatformDirsABCImpl) UserDesktopPath() string {
	return p.X__This.UserDesktopDir()
}

func (p *PlatformDirsABCImpl) UserRuntimePath() string {
	return p.X__This.UserRuntimeDir()
}

func (p *PlatformDirsABCImpl) SiteRuntimePath() string {
	return p.X__This.SiteRuntimeDir()
}

func (p *PlatformDirsABCImpl) IterConfigDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.X__This.UserConfigDir()) {
			return
		}
		if !yield(p.X__This.SiteConfigDir()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterDataDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.X__This.UserDataDir()) {
			return
		}
		if !yield(p.X__This.SiteDataDir()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterCacheDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.X__This.UserCacheDir()) {
			return
		}
		if !yield(p.X__This.SiteCacheDir()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterRuntimeDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.X__This.UserRuntimeDir()) {
			return
		}
		if !yield(p.X__This.SiteRuntimeDir()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterConfigPaths() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.X__This.UserConfigPath()) {
			return
		}
		if !yield(p.X__This.SiteConfigPath()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterDataPaths() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.X__This.UserDataPath()) {
			return
		}
		if !yield(p.X__This.SiteDataPath()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterCachePaths() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.X__This.UserCachePath()) {
			return
		}
		if !yield(p.X__This.SiteCachePath()) {
			return
		}
	}
}

func (p *PlatformDirsABCImpl) IterRuntimePaths() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.X__This.UserRuntimePath()) {
			return
		}
		if !yield(p.X__This.SiteRuntimePath()) {
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
