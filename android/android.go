package android

import (
	"errors"
	"path/filepath"

	"github.com/jcbhmr/go-platformdirs/api"
)

type Android interface {
	api.PlatformDirsABC
}

type AndroidImpl struct {
	api.PlatformDirsABCImpl
}

var _ Android = (*AndroidImpl)(nil)

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) *AndroidImpl {
	this := &AndroidImpl{*api.NewPlatformDirsABC(appname, appauthor, version, roaming, multipath, opinion, ensureExists)}
	this.This = this
	return this
}

func (p *AndroidImpl) UserDataDir() string {
	f, ok := X__AndroidFolder()
	if !ok {
		panic(errors.New("X__AndroidFolder() failed"))
	}
	return p.This.(Android).X__AppendAppNameAndVersion(f, "files")
}

func (p *AndroidImpl) SiteDataDir() string {
	return p.This.(Android).UserDataDir()
}

func (p *AndroidImpl) UserConfigDir() string {
	f, ok := X__AndroidFolder()
	if !ok {
		panic(errors.New("X__AndroidFolder() failed"))
	}
	return p.This.(Android).X__AppendAppNameAndVersion(f, "shared_prefs")
}

func (p *AndroidImpl) SiteConfigDir() string {
	return p.This.(Android).UserConfigDir()
}

func (p *AndroidImpl) UserCacheDir() string {
	f, ok := X__AndroidFolder()
	if !ok {
		panic(errors.New("X__AndroidFolder() failed"))
	}
	return p.This.(Android).X__AppendAppNameAndVersion(f, "cache")
}

func (p *AndroidImpl) SiteCacheDir() string {
	return p.This.(Android).UserCacheDir()
}

func (p *AndroidImpl) UserStateDir() string {
	return p.This.(Android).UserDataDir()
}

func (p *AndroidImpl) UserLogDir() string {
	path := p.This.(Android).UserCacheDir()
	if p.Opinion {
		path = filepath.Join(path, "log")
	}
	return path
}

func (p *AndroidImpl) UserDocumentsDir() string {
	return androidDocumentsFolder()
}

func (p *AndroidImpl) UserDownloadsDir() string {
	return androidDownloadsFolder()
}

func (p *AndroidImpl) UserPicturesDir() string {
	return androidPicturesFolder()
}

func (p *AndroidImpl) UserVideosDir() string {
	return androidVideosFolder()
}

func (p *AndroidImpl) UserMusicDir() string {
	return androidMusicFolder()
}

func (p *AndroidImpl) UserDesktopDir() string {
	return "/storage/emulated/0/Desktop"
}

func (p *AndroidImpl) UserRuntimeDir() string {
	path := p.This.(Android).UserCacheDir()
	if p.Opinion {
		path = filepath.Join(path, "tmp")
	}
	return path
}

func (p *AndroidImpl) SiteRuntimeDir() string {
	return p.This.(Android).UserRuntimeDir()
}

func X__AndroidFolder() (string, bool) {
	panic("not implemented")
}

func androidDocumentsFolder() string {
	panic("not implemented")
}

func androidDownloadsFolder() string {
	panic("not implemented")
}

func androidPicturesFolder() string {
	panic("not implemented")
}

func androidVideosFolder() string {
	panic("not implemented")
}

func androidMusicFolder() string {
	panic("not implemented")
}

