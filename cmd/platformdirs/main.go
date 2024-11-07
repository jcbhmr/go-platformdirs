package main

import (
	"log"

	"github.com/jcbhmr/go-platformdirs"
)

func main() {
	const appName = "MyApp"
	const appAuthor = "MyCompany"

	log.Printf("-- platformdirs %v --", platformdirs.Version)

	log.Printf("-- app dirs (with optional 'version')")
	dirs := platformdirs.New(appName, appAuthor, "1.0", nil, nil, nil, nil)
	log.Printf("UserDataDir(): %v", dirs.UserDataDir())
	log.Printf("UserConfigDir(): %v", dirs.UserConfigDir())
	log.Printf("UserCacheDir(): %v", dirs.UserCacheDir())
	log.Printf("UserStateDir(): %v", dirs.UserStateDir())
	log.Printf("UserLogDir(): %v", dirs.UserLogDir())
	log.Printf("UserDocumentsDir(): %v", dirs.UserDocumentsDir())
	log.Printf("UserDownloadsDir(): %v", dirs.UserDownloadsDir())
	log.Printf("UserPicturesDir(): %v", dirs.UserPicturesDir())
	log.Printf("UserVideosDir(): %v", dirs.UserVideosDir())
	log.Printf("UserMusicDir(): %v", dirs.UserMusicDir())
	log.Printf("UserRuntimeDir(): %v", dirs.UserRuntimeDir())
	log.Printf("SiteDataDir(): %v", dirs.SiteDataDir())
	log.Printf("SiteConfigDir(): %v", dirs.SiteConfigDir())
	log.Printf("SiteCacheDir(): %v", dirs.SiteCacheDir())
	log.Printf("SiteRuntimeDir(): %v", dirs.SiteRuntimeDir())
}
