package windows

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jcbhmr/go-platformdirs/api"
)

func ptr[T any](v T) *T {
	return &v
}

type Windows_X__AppendPartsOptions struct {
	OpinionValue *string
}

type Windows interface {
	api.PlatformDirsABC
	X__AppendParts(path string, options *Windows_X__AppendPartsOptions) string
}

type WindowsImpl struct {
	api.PlatformDirsABCImpl
}

var _ Windows = (*WindowsImpl)(nil)

func New(appname *string, appauthor any, version *string, roaming *bool, multipath *bool, opinion *bool, ensureExists *bool) *WindowsImpl {
	this := &WindowsImpl{*api.NewPlatformDirsABC(appname, appauthor, version, roaming, multipath, opinion, ensureExists)}
	this.X__This = this
	return this
}

func (p *WindowsImpl) UserDataDir() string {
	var const2 string
	if p.X__This.(Windows).Roaming() {
		const2 = "CSIDL_APPDATA"
	} else {
		const2 = "CSIDL_LOCAL_APPDATA"
	}
	path := filepath.Clean(getWinFolder(const2))
	return p.X__This.(Windows).X__AppendParts(path, nil)
}

func (p *WindowsImpl) X__AppendParts(path string, options *Windows_X__AppendPartsOptions) string {
	var opinionValue *string
	if options != nil {
		opinionValue = options.OpinionValue
	} else {
		opinionValue = nil
	}
	params := []string{}
	if appname, ok := p.X__This.(Windows).Appname(); ok && appname != "" {
		if appauthor, ok := p.X__This.(Windows).Appauthor().(bool); !ok || appauthor {
			var author string
			if appauthor, ok := p.X__This.(Windows).Appauthor().(string); ok && appauthor != "" {
				author = appauthor
			} else {
				author = appname
			}
			params = append(params, author)
		}
		params = append(params, appname)
		if opinionValue != nil && *opinionValue != "" {
			params = append(params, *opinionValue)
		}
		if version, ok := p.X__This.(Windows).Version(); ok && version != "" {
			params = append(params, version)
		}
	}
	path = filepath.Join(append([]string{path}, params...)...)
	p.X__This.(Windows).X__OptionallyCreateDirectory(path)
	return path
}

func (p *WindowsImpl) SiteDataDir() string {
	path := filepath.Clean(getWinFolder("CSIDL_COMMON_APPDATA"))
	return p.X__This.(Windows).X__AppendParts(path, nil)
}

func (p *WindowsImpl) UserConfigDir() string {
	return p.X__This.(Windows).UserDataDir()
}

func (p *WindowsImpl) SiteConfigDir() string {
	return p.X__This.(Windows).SiteDataDir()
}

func (p *WindowsImpl) UserCacheDir() string {
	path := filepath.Clean(getWinFolder("CSIDL_LOCAL_APPDATA"))
	return p.X__This.(Windows).X__AppendParts(path, &Windows_X__AppendPartsOptions{OpinionValue: ptr("Cache")})
}

func (p *WindowsImpl) SiteCacheDir() string {
	path := filepath.Clean(getWinFolder("CSIDL_COMMON_APPDATA"))
	return p.X__This.(Windows).X__AppendParts(path, &Windows_X__AppendPartsOptions{OpinionValue: ptr("Cache")})
}

func (p *WindowsImpl) UserStateDir() string {
	return p.X__This.(Windows).UserDataDir()
}

func (p *WindowsImpl) UserLogDir() string {
	path := p.X__This.(Windows).UserDataDir()
	if p.X__This.(Windows).Opinion() {
		path = filepath.Join(path, "Logs")
		p.X__This.(Windows).X__OptionallyCreateDirectory(path)
	}
	return path
}

func (p *WindowsImpl) UserDocumentsDir() string {
	return filepath.Clean(getWinFolder("CSIDL_PERSONAL"))
}

func (p *WindowsImpl) UserDownloadsDir() string {
	return filepath.Clean(getWinFolder("CSIDL_DOWNLOADS"))
}

func (p *WindowsImpl) UserPicturesDir() string {
	return filepath.Clean(getWinFolder("CSIDL_MYPICTURES"))
}

func (p *WindowsImpl) UserVideosDir() string {
	return filepath.Clean(getWinFolder("CSIDL_MYVIDEO"))
}

func (p *WindowsImpl) UserMusicDir() string {
	return filepath.Clean(getWinFolder("CSIDL_MYMUSIC"))
}

func (p *WindowsImpl) UserDesktopDir() string {
	return filepath.Clean(getWinFolder("CSIDL_DESKTOPDIRECTORY"))
}

func (p *WindowsImpl) UserRuntimeDir() string {
	path := filepath.Clean(filepath.Join(getWinFolder("CSIDL_LOCAL_APPDATA"), "Temp"))
	return p.X__This.(Windows).X__AppendParts(path, nil)
}

func (p *WindowsImpl) SiteRuntimeDir() string {
	return p.X__This.(Windows).UserRuntimeDir()
}

func getWinFolderFromEnvVars(csidlName string) string {
	result, ok := getWinFolderIfCSIDLNameNotEnvVar(csidlName)
	if ok {
		return result
	}

	envVarName, ok := map[string]string{
		"CSIDL_APPDATA":        "APPDATA",
		"CSIDL_COMMON_APPDATA": "ALLUSERSPROFILE",
		"CSIDL_LOCAL_APPDATA":  "LOCALAPPDATA",
	}[csidlName]
	if !ok {
		panic(fmt.Errorf("unknown CSIDL name: %s", csidlName))
	}
	result, ok = os.LookupEnv(envVarName)
	if !ok {
		panic(fmt.Errorf("unset environment variable: %s", envVarName))
	}
	return result
}

func getWinFolderIfCSIDLNameNotEnvVar(csidlName string) (string, bool) {
	if csidlName == "CSIDL_PERSONAL" {
		return filepath.Join(filepath.Clean(os.Getenv("USERPROFILE")), "Documents"), true
	} else if csidlName == "CSIDL_DOWNLOADS" {
		return filepath.Join(filepath.Clean(os.Getenv("USERPROFILE")), "Downloads"), true
	} else if csidlName == "CSIDL_MYPICTURES" {
		return filepath.Join(filepath.Clean(os.Getenv("USERPROFILE")), "Pictures"), true
	} else if csidlName == "CSIDL_MYVIDEO" {
		return filepath.Join(filepath.Clean(os.Getenv("USERPROFILE")), "Videos"), true
	} else if csidlName == "CSIDL_MYMUSIC" {
		return filepath.Join(filepath.Clean(os.Getenv("USERPROFILE")), "Music"), true
	}
	return "", false
}

func pickGetWinFolder() func(string) string {
	return getWinFolderFromEnvVars
}

var getWinFolder = pickGetWinFolder()
