# Platform details

platformdirs auto-detects the current platform and returns the correct directory paths. This page describes the default paths for each platform and any platform-specific behaviour.

All examples below assume AppAuthor: "SuperApp" and AppAuthor: "Acme" unless stated otherwise.

## User directories

These are user-specific (and, generally, user-writeable) directories.

### UserDataDir

 - Linux: `~/.local/share/SuperApp`
 - macOS: `~/Library/Application Support/SuperApp`
 - Windows: `C:\Users\<User>\AppData\Local\Acme\SuperApp`
 - Android: `/data/data/<pkg>/files/SuperApp`

### UserConfigDir

 - Linux: `~/.config/SuperApp`
 - macOS: `~/Library/Application Support/SuperApp`
 - Windows: `C:\Users\<User>\AppData\Local\Acme\SuperApp`
 - Android: `/data/data/<pkg>/shared_prefs/SuperApp`

### UserCacheDir

 - Linux: `~/.cache/SuperApp`
 - macOS: `~/Library/Caches/SuperApp`
 - Windows: `C:\Users\<User>\AppData\Local\Acme\SuperApp\Cache`
 - Android: `/data/data/<pkg>/cache/SuperApp`

### UserStateDir

 - Linux: `~/.local/state/SuperApp`
 - macOS: `~/Library/Application Support/SuperApp`
 - Windows: `C:\Users\<User>\AppData\Local\Acme\SuperApp`
 - Android: `/data/data/<pkg>/files/SuperApp`

### UserLogDir

 - Linux: `~/.local/state/SuperApp/log`
 - macOS: `~/Library/Logs/SuperApp`
 - Windows: `C:\Users\<User>\AppData\Local\Acme\SuperApp\Logs`
 - Android: `/data/data/<pkg>/cache/SuperApp/log`

### UserRuntimeDir

 - Linux: `/run/user/<uid>/SuperApp`
 - macOS: `~/Library/Caches/TemporaryItems/SuperApp`
 - Windows: `C:\Users\<User>\AppData\Local\Temp\Acme\SuperApp`
 - Android: `/data/data/<pkg>/cache/SuperApp/tmp`

### UserApplicationsDir

 - Linux: `~/.local/share/applications`
 - macOS: `~/Applications`
 - Windows: `C:\Users\<User>\AppData\Roaming\Microsoft\Windows\Start Menu\Programs`
 - Android: Same as `user_data_dir`

Note: This property does not append appname or version. It returns the shared applications directory where .desktop files (Linux), app bundles (macOS), or Start Menu shortcuts (Windows) are placed.

### UserBinDir

 - Linux: `~/.local/bin`
 - macOS: `~/.local/bin`
 - Windows: `C:\Users\<User>\AppData\Local\Programs`
 - Android: `/data/data/<pkg>/files/bin`

Note: This property does not append appname or version. It returns the directory where user-installed executables and scripts are placed.

### UserDocumentsDir

 - Linux: `~/Documents`
 - macOS: `~/Documents`
 - Windows: `C:\Users\<User>\Documents`
 - Android: `/storage/emulated/0/Documents`

### UserDownloadsDir

 - Linux: `~/Downloads`
 - macOS: `~/Downloads`
 - Windows: `C:\Users\<User>\Downloads`
 - Android: `/storage/emulated/0/Downloads`

### UserPicturesDir

 - Linux: `~/Pictures`
 - macOS: `~/Pictures`
 - Windows: `C:\Users\<User>\Pictures`
 - Android: `/storage/emulated/0/Pictures`

### UserVideosDir

 - Linux: `~/Videos`
 - macOS: `~/Movies`
 - Windows: `C:\Users\<User>\Videos`
 - Android: `/storage/emulated/0/DCIM/Camera`

### UserMusicDir

 - Linux: `~/Music`
 - macOS: `~/Music`
 - Windows: `C:\Users\<User>\Music`
 - Android: `/storage/emulated/0/Music`

### UserDesktopDir

 - Linux: `~/Desktop`
 - macOS: `~/Desktop`
 - Windows: `C:\Users\<User>\Desktop`
 - Android: `/storage/emulated/0/Desktop`

# Shared directories

These are system-wide (and, generally, read-only) directories.

# SiteDataDir

 - Linux: `/usr/local/share/SuperApp`
 - macOS: `/Library/Application Support/SuperApp`
 - Windows: `C:\ProgramData\Acme\SuperApp`
 - Android: `/data/data/<pkg>/files/SuperApp`

# SiteConfigDir

 - Linux: `/etc/xdg/SuperApp`
 - macOS: `/Library/Application Support/SuperApp`
 - Windows: `C:\ProgramData\Acme\SuperApp`
 - Android: `/data/data/<pkg>/shared_prefs/SuperApp`

# SiteCacheDir

 - Linux: `/var/cache/SuperApp`
 - macOS: `/Library/Caches/SuperApp`
 - Windows: `C:\ProgramData\Acme\SuperApp\Cache`
 - Android: `/data/data/<pkg>/cache/SuperApp`

# SiteStateDir

 - Linux: `/var/lib/SuperApp`
 - macOS: `/Library/Application Support/SuperApp`
 - Windows: `C:\ProgramData\Acme\SuperApp`
 - Android: `/data/data/<pkg>/files/SuperApp`

# SiteLogDir

 - Linux: `/var/log/SuperApp`
 - macOS: `/Library/Logs/SuperApp`
 - Windows: `C:\ProgramData\Acme\SuperApp\Logs`
 - Android: `/data/data/<pkg>/cache/SuperApp/log`

# SiteRuntimeDir

 - Linux: `/run/SuperApp`
 - macOS: `~/Library/Caches/TemporaryItems/SuperApp`
 - Windows: `C:\Users\<User>\AppData\Local\Temp\Acme\SuperApp`
 - Android: `/data/data/<pkg>/cache/SuperApp/tmp`

# SiteApplicationsDir

 - Linux: `/usr/share/applications`
 - macOS: `/Applications`
 - Windows: `C:\ProgramData\Microsoft\Windows\Start Menu\Programs`
 - Android: Same as `user_applications_dir`

Note: This property does not append appname or version. It returns the system-wide applications directory where .desktop files (Linux), app bundles (macOS), or Start Menu shortcuts (Windows) are installed for all users.

# SiteBinDir

 - Linux: `/usr/local/bin`
 - macOS: `/usr/local/bin`
 - Windows: `C:\ProgramData\bin`
 - Android: Same as `user_bin_dir`

Note: This property does not append appname or version. It returns the directory where system-wide executables and scripts are placed. On Unix/Linux, this follows the FHS 3.0 standard for locally installed software. On Windows, it mirrors the site_data_dir pattern using %ProgramData%, following the precedent set by Chocolatey.

See also: For platform-specific conventions and behavior details, see Understanding platformdirs.
