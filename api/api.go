package api

import (
	"iter"
)

type PlatformDirsABC interface {
	Appname() (string, bool)
	Appauthor() (any, bool)
	Version() (string, bool)
	Roaming() bool
	Multipath() bool
	Opinion() bool
	EnsureExists() bool
	UserDataDir() (string, error)
	SiteDataDir() (string, error)
	UserConfigDir() (string, error)
	SiteConfigDir() (string, error)
	UserCacheDir() (string, error)
	SiteCacheDir() (string, error)
	UserStateDir() (string, error)
	UserLogDir() (string, error)
	UserDocumentsDir() (string, error)
	UserDownloadsDir() (string, error)
	UserPicturesDir() (string, error)
	UserVideosDir() (string, error)
	UserMusicDir() (string, error)
	UserDesktopDir() (string, error)
	UserRuntimeDir() (string, error)
	SiteRuntimeDir() (string, error)
	UserDataPath() (string, error)
	SiteDataPath() (string, error)
	UserConfigPath() (string, error)
	SiteConfigPath() (string, error)
	UserCachePath() (string, error)
	SiteCachePath() (string, error)
	UserStatePath() (string, error)
	UserLogPath() (string, error)
	UserDocumentsPath() (string, error)
	UserDownloadsPath() (string, error)
	UserPicturesPath() (string, error)
	UserVideosPath() (string, error)
	UserMusicPath() (string, error)
	UserDesktopPath() (string, error)
	UserRuntimePath() (string, error)
	SiteRuntimePath() (string, error)
	IterConfigDirs() iter.Seq[string]
	IterDataDirs() iter.Seq[string]
	IterCacheDirs() iter.Seq[string]
	IterRuntimeDirs() iter.Seq[string]
	IterConfigPaths() iter.Seq[string]
	IterDataPaths() iter.Seq[string]
	IterCachePaths() iter.Seq[string]
	IterRuntimePaths() iter.Seq[string]
}
