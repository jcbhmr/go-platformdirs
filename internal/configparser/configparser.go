package configparser

import (
	"regexp"
	"strings"
)

// Stub of to https://docs.python.org/3/library/configparser.html.
// Must parse files like https://github.com/search?q=path%3Auser-dirs.dirs&type=code.

/*
# comment
[section]
key = value
*/

/*
# skeleton ~/.config/user-dirs.dirs
XDG_DESKTOP_DIR="$HOME/.local/share/Desktop"
XDG_DOCUMENTS_DIR="$HOME/Dropbox"
XDG_DOWNLOAD_DIR="$HOME/Downloads"
XDG_MUSIC_DIR="$HOME/Music"
XDG_PICTURES_DIR="$HOME/Pictures"
*/

type ConfigParser struct {
	data map[string]map[string]string
}

func New() *ConfigParser {
	return &ConfigParser{}
}

func (c *ConfigParser) ReadString(input string) {
	data := map[string]map[string]string{}
	sectionName := ""
	for _, line := range strings.FieldsFunc(input, func(c rune) bool { return c == '\n' || c == '\r' }) {
		if strings.HasPrefix(line, "#") {
			continue
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		match := regexp.MustCompile(`^\[(.*)\]$`).FindStringSubmatch(line)
		if len(match) > 0 {
			sectionName = match[1]
			continue
		}

		if _, ok := data[sectionName]; !ok {
			data[sectionName] = map[string]string{}
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		data[sectionName][key] = value
	}
}

func (c *ConfigParser) Get(key string) (map[string]string, bool) {
	v, ok := c.data[key]
	return v, ok
}
