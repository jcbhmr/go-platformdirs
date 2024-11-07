package api

import (
	"os"
	"path/filepath"
)

type PlatformDirsABC struct {
	Appname2      *string
	Appauthor2    any
	Version2      *string
	Roaming2      bool
	Multipath2    bool
	Opinion2      bool
	EnsureExists2 bool
}

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) PlatformDirsABC {
	var appauthor2 any
	if appauthor == nil {
		appauthor2 = nil
	} else if v, ok := appauthor.(string); ok {
		appauthor2 = v
	} else if v, ok := appauthor.(bool); ok && !v {
		appauthor2 = false
	} else {
		panic("appauthor must be string, false, or nil")
	}
	roaming2 := false
	if roaming != nil {
		roaming2 = *roaming
	}
	multipath2 := false
	if multipath != nil {
		multipath2 = *multipath
	}
	opinion2 := true
	if opinion != nil {
		opinion2 = *opinion
	}
	ensureExists2 := false
	if ensureExists != nil {
		ensureExists2 = *ensureExists
	}
	return PlatformDirsABC{
		Appname2:      appname,
		Appauthor2:    appauthor2,
		Version2:      version,
		Roaming2:      roaming2,
		Multipath2:    multipath2,
		Opinion2:      opinion2,
		EnsureExists2: ensureExists2,
	}
}

func (p *PlatformDirsABC) AppendAppNameAndVersion(base ...string) (string, error) {
	params := base[1:]
	if p.Appname2 != nil && *p.Appname2 != "" {
		params = append(params, *p.Appname2)
		if p.Version2 != nil && *p.Version2 != "" {
			params = append(params, *p.Version2)
		}
	}
	joinArgs := make([]string, 0, 1+len(params))
	joinArgs = append(joinArgs, base[0])
	joinArgs = append(joinArgs, params...)
	path := filepath.Join(joinArgs...)
	err := p.OptionallyCreateDirectory(path)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (p *PlatformDirsABC) OptionallyCreateDirectory(path string) error {
	if p.EnsureExists2 {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func (p *PlatformDirsABC) FirstItemAsPathIfMultipath(directory string) string {
	if p.Multipath2 {
		directory, _ = filepath.Split(directory)
	}
	return directory
}

func (p *PlatformDirsABC) UserDataDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) SiteDataDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserConfigDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) SiteConfigDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserCacheDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) SiteCacheDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserStateDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserLogDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserDocumentsDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserDownloadsDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserPicturesDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserVideosDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserMusicDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserDesktopDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserRuntimeDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) SiteRuntimeDir() (string, error) {
	panic("abstract method")
}

func (p *PlatformDirsABC) UserDataPath() (string, error) {
	return p.UserDataDir()
}

func (p *PlatformDirsABC) SiteDataPath() (string, error) {
	return p.SiteDataDir()
}

func (p *PlatformDirsABC) UserConfigPath() (string, error) {
	return p.UserConfigDir()
}

func (p *PlatformDirsABC) SiteConfigPath() (string, error) {
	return p.SiteConfigDir()
}

func (p *PlatformDirsABC) UserCachePath() (string, error) {
	return p.UserCacheDir()
}

func (p *PlatformDirsABC) SiteCachePath() (string, error) {
	return p.SiteCacheDir()
}

func (p *PlatformDirsABC) UserStatePath() (string, error) {
	return p.UserStateDir()
}

func (p *PlatformDirsABC) UserLogPath() (string, error) {
	return p.UserLogDir()
}

func (p *PlatformDirsABC) UserDocumentsPath() (string, error) {
	return p.UserDocumentsDir()
}

func (p *PlatformDirsABC) UserDownloadsPath() (string, error) {
	return p.UserDownloadsDir()
}

func (p *PlatformDirsABC) UserPicturesPath() (string, error) {
	return p.UserPicturesDir()
}

func (p *PlatformDirsABC) UserVideosPath() (string, error) {
	return p.UserVideosDir()
}

func (p *PlatformDirsABC) UserMusicPath() (string, error) {
	return p.UserMusicDir()
}

