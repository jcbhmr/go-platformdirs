package windows

import (
	"path/filepath"

	"github.com/jcbhmr/go-platformdirs/api"
	internalapi "github.com/jcbhmr/go-platformdirs/internal/api"
	"github.com/samber/mo"
)

type Windows interface {
	api.PlatformDirsABC
}

type windowsImpl struct {
	internalapi.PlatformDirsABCImpl
}

func (w *windowsImpl) UserDataDir() {
	var const2 string
	if w.Roaming {
		const2 = "CSIDL_APPDATA"
	} else {
		const2 = "CSIDL_LOCAL_APPDATA"
	}
	path, err := getWinFolder(const2)
	if err != nil {
		return "", err
	}
	path = filepath.Clean(path)
	return w.appendParts(path)
}

func (w *windowsImpl) appendParts(path string, options struct{ opinionValue mo.Option[string] }) {
	params := []string{}
	if appname, ok := w.Appname; ok && appname != "" {

	}
}
