package unix

import (
	"errors"
	"fmt"
	"iter"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/jcbhmr/go-platformdirs/api"
	"github.com/jcbhmr/go-platformdirs/internal/configparser"
)

func mustUserHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("os.UserHomeDir() failed: %w", err))
	}
	return homeDir
}

var getuid func() int

func init() {
	if runtime.GOOS == "windows" {
		getuid = func() int {
			panic(errors.New("should only be used on Unix"))
		}
	} else {
		getuid = os.Getuid
	}
}

type Unix interface {
	api.PlatformDirsABC
	X__SiteDataDirs() []string
	X__SiteConfigDirs() []string
}

type UnixImpl struct {
	api.PlatformDirsABCImpl
}

var _ Unix = (*UnixImpl)(nil)

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) *UnixImpl {
	this := &UnixImpl{*api.NewPlatformDirsABC(appname, appauthor, version, roaming, multipath, opinion, ensureExists)}
	this.This = this
	return this
}

func (p *UnixImpl) UserDataDir() string {
	path := os.Getenv("XDG_DATA_HOME")
	if strings.TrimSpace(path) == "" {
		path = filepath.Join(mustUserHomeDir(), ".local/share")
	}
	return p.This.(Unix).X__AppendAppNameAndVersion(path)
}

func (p *UnixImpl) X__SiteDataDirs() []string {
	path := os.Getenv("XDG_DATA_DIRS")
	if strings.TrimSpace(path) == "" {
		path = fmt.Sprintf("/usr/local/share%s/usr/share", string(filepath.ListSeparator))
	}
	r := []string{}
	for _, p2 := range filepath.SplitList(path) {
		r = append(r, p.This.(Unix).X__AppendAppNameAndVersion(p2))
	}
	return r
}

func (p *UnixImpl) SiteDataDir() string {
	dirs := p.This.(Unix).X__SiteDataDirs()
	if !p.Multipath {
		return dirs[0]
	}
	return strings.Join(dirs, string(filepath.ListSeparator))
}

func (p *UnixImpl) UserConfigDir() string {
	path := os.Getenv("XDG_CONFIG_HOME")
	if strings.TrimSpace(path) == "" {
		path = filepath.Join(mustUserHomeDir(), ".config")
	}
	return p.This.(Unix).X__AppendAppNameAndVersion(path)
}

func (p *UnixImpl) X__SiteConfigDirs() []string {
	path := os.Getenv("XDG_CONFIG_DIRS")
	if strings.TrimSpace(path) == "" {
		path = "/etc/xdg"
	}
	r := []string{}
	for _, p2 := range filepath.SplitList(path) {
		r = append(r, p.This.(Unix).X__AppendAppNameAndVersion(p2))
	}
	return r
}

func (p *UnixImpl) SiteConfigDir() string {
	dirs := p.This.(Unix).X__SiteConfigDirs()
	if !p.Multipath {
		return dirs[0]
	}
	return strings.Join(dirs, string(filepath.ListSeparator))
}

func (p *UnixImpl) UserCacheDir() string {
	path := os.Getenv("XDG_CACHE_HOME")
	if strings.TrimSpace(path) == "" {
		path = filepath.Join(mustUserHomeDir(), ".cache")
	}
	return p.This.(Unix).X__AppendAppNameAndVersion(path)
}

func (p *UnixImpl) SiteCacheDir() string {
	return p.This.(Unix).X__AppendAppNameAndVersion("/var/cache")
}

func (p *UnixImpl) UserStateDir() string {
	path := os.Getenv("XDG_STATE_HOME")
	if strings.TrimSpace(path) == "" {
		path = filepath.Join(mustUserHomeDir(), ".local/state")
	}
	return p.This.(Unix).X__AppendAppNameAndVersion(path)
}

func (p *UnixImpl) UserLogDir() string {
	path := p.This.(Unix).UserStateDir()
	if p.Opinion {
		path = filepath.Join(path, "log")
		p.This.(Unix).X__OptionallyCreateDirectory(path)
	}
	return path
}

func (p *UnixImpl) UserDocumentsDir() string {
	return getUserMediaDir("XDG_DOCUMENTS_DIR", "~/Documents")
}

func (p *UnixImpl) UserDownloadsDir() string {
	return getUserMediaDir("XDG_DOWNLOAD_DIR", "~/Downloads")
}

func (p *UnixImpl) UserPicturesDir() string {
	return getUserMediaDir("XDG_PICTURES_DIR", "~/Pictures")
}

func (p *UnixImpl) UserVideosDir() string {
	return getUserMediaDir("XDG_VIDEOS_DIR", "~/Videos")
}

func (p *UnixImpl) UserMusicDir() string {
	return getUserMediaDir("XDG_MUSIC_DIR", "~/Music")
}

func (p *UnixImpl) UserDesktopDir() string {
	return getUserMediaDir("XDG_DESKTOP_DIR", "~/Desktop")
}

func (p *UnixImpl) UserRuntimeDir() string {
	path := os.Getenv("XDG_RUNTIME_DIR")
	if strings.TrimSpace(path) == "" {
		if runtime.GOOS == "freebsd" || runtime.GOOS == "openbsd" || runtime.GOOS == "netbsd" {
			path = fmt.Sprintf("/var/run/user/%d", getuid())
			if _, err := os.Stat(path); err != nil {
				path = fmt.Sprintf("/tmp/runtime-%d", getuid())
			}
		} else {
			path = fmt.Sprintf("/run/user/%d", getuid())
		}
	}
	return p.This.(Unix).X__AppendAppNameAndVersion(path)
}

func (p *UnixImpl) SiteRuntimeDir() string {
	path := os.Getenv("XDG_RUNTIME_DIR")
	if strings.TrimSpace(path) == "" {
		if runtime.GOOS == "freebsd" || runtime.GOOS == "openbsd" || runtime.GOOS == "netbsd" {
			path = "/var/run"
		} else {
			path = "/run"
		}
	}
	return p.This.(Unix).X__AppendAppNameAndVersion(path)
}

func (p *UnixImpl) SiteDataPath() string {
	return p.This.(Unix).X__FirstItemAsPathIfMultipath(p.This.(Unix).SiteDataDir())
}

func (p *UnixImpl) SiteConfigPath() string {
	return p.This.(Unix).X__FirstItemAsPathIfMultipath(p.This.(Unix).SiteConfigDir())
}

func (p *UnixImpl) SiteCachePath() string {
	return p.This.(Unix).X__FirstItemAsPathIfMultipath(p.This.(Unix).SiteCacheDir())
}

func (p *UnixImpl) IterConfigDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.(Unix).UserConfigDir()) {
			return
		}
		for _, x := range p.This.(Unix).X__SiteConfigDirs() {
			if !yield(x) {
				return
			}
		}
	}
}

func (p *UnixImpl) IterDataDirs() iter.Seq[string] {
	return func(yield func(string) bool) {
		if !yield(p.This.(Unix).UserDataDir()) {
			return
		}
		for _, x := range p.This.(Unix).X__SiteDataDirs() {
			if !yield(x) {
				return
			}
		}
	}
}

func getUserMediaDir(envVar string, fallbackTildePath string) string {
	mediaDir, ok := getUserDirsFolder(envVar)
	if !ok {
		mediaDir = strings.TrimSpace(os.Getenv(envVar))
		if mediaDir == "" {
			mediaDir = regexp.MustCompile(`^~`).ReplaceAllString(fallbackTildePath, mustUserHomeDir())
		}
	}
	return mediaDir
}

func getUserDirsFolder(envVar string) (string, bool) {
	userDirsConfigPath := filepath.Join(New(nil, nil, nil, nil, nil, nil, nil).UserConfigDir(), "user-dirs.dirs")
	if _, err := os.Stat(userDirsConfigPath); err == nil {
		parser := configparser.New()

		read, err := os.ReadFile(userDirsConfigPath)
		if err != nil {
			panic(fmt.Errorf("os.ReadFile() failed: %w", err))
		}
		parser.ReadString("[top]\n" + string(read))

		top, ok := parser.Get("top")
		if !ok {
			return "", false
		} else if _, ok := top[envVar]; !ok {
			return "", false
		}

		path := strings.Trim(top[envVar], `"`)
		return strings.Replace(path, "$HOME", mustUserHomeDir(), 1), true
	}
	return "", false
}
