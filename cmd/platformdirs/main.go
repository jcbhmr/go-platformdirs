package main

import (
	"fmt"
	"log"

	"github.com/jcbhmr/go-platformdirs"
)

func ptr[T any](v T) *T {
	return &v
}

/*
PROPS = (
    "user_data_dir",
    "user_config_dir",
    "user_cache_dir",
    "user_state_dir",
    "user_log_dir",
    "user_documents_dir",
    "user_downloads_dir",
    "user_pictures_dir",
    "user_videos_dir",
    "user_music_dir",
    "user_runtime_dir",
    "site_data_dir",
    "site_config_dir",
    "site_cache_dir",
    "site_runtime_dir",
)
*/

func main() {
	log.SetFlags(0)

	appName := "MyApp"
	appAuthor := "MyCompany"

	fmt.Printf("-- platformdirs %s --\n", platformdirs.Version)

	fmt.Println("-- app dirs (with optional 'version')")
	dirs := platformdirs.New(&appName, appAuthor, ptr("1.0"), nil, nil, nil, nil)
	{
		fmt.Printf("%s: %s\n", "user_data_dir", dirs.UserDataDir())
		fmt.Printf("%s: %s\n", "user_config_dir", dirs.UserConfigDir())
		fmt.Printf("%s: %s\n", "user_cache_dir", dirs.UserCacheDir())
		fmt.Printf("%s: %s\n", "user_state_dir", dirs.UserStateDir())
		fmt.Printf("%s: %s\n", "user_log_dir", dirs.UserLogDir())
		fmt.Printf("%s: %s\n", "user_documents_dir", dirs.UserDocumentsDir())
		fmt.Printf("%s: %s\n", "user_downloads_dir", dirs.UserDownloadsDir())
		fmt.Printf("%s: %s\n", "user_pictures_dir", dirs.UserPicturesDir())
		fmt.Printf("%s: %s\n", "user_videos_dir", dirs.UserVideosDir())
		fmt.Printf("%s: %s\n", "user_music_dir", dirs.UserMusicDir())
		fmt.Printf("%s: %s\n", "user_runtime_dir", dirs.UserRuntimeDir())
		fmt.Printf("%s: %s\n", "site_data_dir", dirs.SiteDataDir())
		fmt.Printf("%s: %s\n", "site_config_dir", dirs.SiteConfigDir())
		fmt.Printf("%s: %s\n", "site_cache_dir", dirs.SiteCacheDir())
		fmt.Printf("%s: %s\n", "site_runtime_dir", dirs.SiteRuntimeDir())
	}

	fmt.Println("\n-- app dirs (without optional 'version')")
	dirs = platformdirs.New(&appName, appAuthor, nil, nil, nil, nil, nil)
	{
		fmt.Printf("%s: %s\n", "user_data_dir", dirs.UserDataDir())
		fmt.Printf("%s: %s\n", "user_config_dir", dirs.UserConfigDir())
		fmt.Printf("%s: %s\n", "user_cache_dir", dirs.UserCacheDir())
		fmt.Printf("%s: %s\n", "user_state_dir", dirs.UserStateDir())
		fmt.Printf("%s: %s\n", "user_log_dir", dirs.UserLogDir())
		fmt.Printf("%s: %s\n", "user_documents_dir", dirs.UserDocumentsDir())
		fmt.Printf("%s: %s\n", "user_downloads_dir", dirs.UserDownloadsDir())
		fmt.Printf("%s: %s\n", "user_pictures_dir", dirs.UserPicturesDir())
		fmt.Printf("%s: %s\n", "user_videos_dir", dirs.UserVideosDir())
		fmt.Printf("%s: %s\n", "user_music_dir", dirs.UserMusicDir())
		fmt.Printf("%s: %s\n", "user_runtime_dir", dirs.UserRuntimeDir())
		fmt.Printf("%s: %s\n", "site_data_dir", dirs.SiteDataDir())
		fmt.Printf("%s: %s\n", "site_config_dir", dirs.SiteConfigDir())
		fmt.Printf("%s: %s\n", "site_cache_dir", dirs.SiteCacheDir())
		fmt.Printf("%s: %s\n", "site_runtime_dir", dirs.SiteRuntimeDir())
	}

	fmt.Println("\n-- app dirs (without optional 'appauthor')")
	dirs = platformdirs.New(&appName, nil, nil, nil, nil, nil, nil)
	{
		fmt.Printf("%s: %s\n", "user_data_dir", dirs.UserDataDir())
		fmt.Printf("%s: %s\n", "user_config_dir", dirs.UserConfigDir())
		fmt.Printf("%s: %s\n", "user_cache_dir", dirs.UserCacheDir())
		fmt.Printf("%s: %s\n", "user_state_dir", dirs.UserStateDir())
		fmt.Printf("%s: %s\n", "user_log_dir", dirs.UserLogDir())
		fmt.Printf("%s: %s\n", "user_documents_dir", dirs.UserDocumentsDir())
		fmt.Printf("%s: %s\n", "user_downloads_dir", dirs.UserDownloadsDir())
		fmt.Printf("%s: %s\n", "user_pictures_dir", dirs.UserPicturesDir())
		fmt.Printf("%s: %s\n", "user_videos_dir", dirs.UserVideosDir())
		fmt.Printf("%s: %s\n", "user_music_dir", dirs.UserMusicDir())
		fmt.Printf("%s: %s\n", "user_runtime_dir", dirs.UserRuntimeDir())
		fmt.Printf("%s: %s\n", "site_data_dir", dirs.SiteDataDir())
		fmt.Printf("%s: %s\n", "site_config_dir", dirs.SiteConfigDir())
		fmt.Printf("%s: %s\n", "site_cache_dir", dirs.SiteCacheDir())
		fmt.Printf("%s: %s\n", "site_runtime_dir", dirs.SiteRuntimeDir())
	}

	fmt.Println("\n-- app dirs (with disabled 'appauthor')")
	dirs = platformdirs.New(&appName, false, nil, nil, nil, nil, nil)
	{
		fmt.Printf("%s: %s\n", "user_data_dir", dirs.UserDataDir())
		fmt.Printf("%s: %s\n", "user_config_dir", dirs.UserConfigDir())
		fmt.Printf("%s: %s\n", "user_cache_dir", dirs.UserCacheDir())
		fmt.Printf("%s: %s\n", "user_state_dir", dirs.UserStateDir())
		fmt.Printf("%s: %s\n", "user_log_dir", dirs.UserLogDir())
		fmt.Printf("%s: %s\n", "user_documents_dir", dirs.UserDocumentsDir())
		fmt.Printf("%s: %s\n", "user_downloads_dir", dirs.UserDownloadsDir())
		fmt.Printf("%s: %s\n", "user_pictures_dir", dirs.UserPicturesDir())
		fmt.Printf("%s: %s\n", "user_videos_dir", dirs.UserVideosDir())
		fmt.Printf("%s: %s\n", "user_music_dir", dirs.UserMusicDir())
		fmt.Printf("%s: %s\n", "user_runtime_dir", dirs.UserRuntimeDir())
		fmt.Printf("%s: %s\n", "site_data_dir", dirs.SiteDataDir())
		fmt.Printf("%s: %s\n", "site_config_dir", dirs.SiteConfigDir())
		fmt.Printf("%s: %s\n", "site_cache_dir", dirs.SiteCacheDir())
		fmt.Printf("%s: %s\n", "site_runtime_dir", dirs.SiteRuntimeDir())
	}
}
