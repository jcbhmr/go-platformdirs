# platformdirs for Go

<div align=center>
<table><td>

```go
TODO
```

</table>
</div>

## Installation

```sh
go get go.jcbhmr.com/platformdirs
```

## Usage

```go
dirs := platformdirs.NewPlatformDirs("MyApp", "MyCompany")
dirs.UserDataDir()      // /home/octocat/.local/share/MyApp
dirs.UserConfigDir()    // /home/octocat/.config/MyApp
dirs.UserCacheDir()     // /home/octocat/.cache/MyApp
dirs.UserStateDir()     // /home/octocat/.local/state/MyApp
dirs.UserLogDir()       // /home/octocat/.local/state/MyApp/log
dirs.UserDocumentsDir() // /home/octocat/Documents
dirs.UserDownloadsDir() // /home/octocat/Downloads
dirs.UserRuntimeDir()   // /run/user/1234/MyApp
```

For `Path` objects instead of strings:

```go
dirs := platformdirs.NewPlatformDirs("MyApp", "MyCompany")
dirs.UserDataPath()     // TODO
dirs.UserConfigPath()   // TODO
```

Convenience functions for quick access:

```go
platformdirs.UserDataDir("MyApp", "MyCompany")      // returns string
platformdirs.UserConfigPath("MyApp", "MyCompany")   // TODO
```


### Directory types

- **Data:** Persistent application data (user_data_dir, site_data_dir)
- **Config:** Configuration files and settings (user_config_dir, site_config_dir)
- **Cache:** Cached data that can be regenerated (user_cache_dir, site_cache_dir)
- **State:** Non-essential runtime state like window positions (user_state_dir, site_state_dir)
- **Logs:** Log files (user_log_dir, site_log_dir)
- **Runtime:** Runtime files like sockets and PIDs (user_runtime_dir, site_runtime_dir)

Each type has both user_* (per-user, writable) and site_* (system-wide, read-only for users) variants.

### Documentation

Full documentation is available at [pkg.go.dev/go.jcbhmr.com/platformdirs/v4](https://pkg.go.dev/go.jcbhmr.com/platformdirs/v4):

- Getting started tutorial -- learn core concepts through real-world examples
- How-to guides -- recipes for common tasks and platform-specific tips
- API reference -- complete list of functions and classes
- Platform details -- default paths for each operating system

Contributions are welcome! See CONTRIBUTING.md for details.

## Development


