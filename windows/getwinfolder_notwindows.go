//go:build !windows

package windows

import (
	"fmt"
	"os"
	"path/filepath"
)

func getWinFolderFromEnvVars(csidlName string) (string, error) {
	result, ok := getWinFolderIfCSIDLNameNotEnvVar(csidlName)
	if ok {
		return result, nil
	}

	envVarName, ok := map[string]string{
		"CSIDL_APPDATA":        "APPDATA",
		"CSIDL_COMMON_APPDATA": "ALLUSERSPROFILE",
		"CSIDL_LOCAL_APPDATA":  "LOCALAPPDATA",
	}[csidlName]
	if !ok {
		return "", fmt.Errorf("Unknown CSIDL name: %v", csidlName)
	}
	result, ok = os.LookupEnv(envVarName)
	if !ok {
		return "", fmt.Errorf("Unset environment variable: %v", envVarName)
	}
	return result, nil
}

func getWinFolderIfCSIDLNameNotEnvVar(csidlName string) (string, bool) {
	if csidlName == "CSIDL_PERSONAL" {
		return filepath.Join(filepath.Clean(os.Getenv("USERPROFILE")), "Documents"), true
	}

	if csidlName == "CSIDL_DOWNLOADS" {
		return filepath.Join(filepath.Clean(os.Getenv("USERPROFILE")), "Downloads"), true
	}

	if csidlName == "CSIDL_MYPICTURES" {
		return filepath.Join(filepath.Clean(os.Getenv("USERPROFILE")), "Pictures"), true
	}

	if csidlName == "CSIDL_MYMUSIC" {
		return filepath.Join(filepath.Clean(os.Getenv("USERPROFILE")), "Music"), true
	}

	return "", false
}

func getWinFolder(csidlName string) (string, error) {
	return getWinFolderFromEnvVars(csidlName)
}
