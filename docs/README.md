# platformdirs
platformdirs is a Python library for determining platform-specific system directories. Whether you need user data, configuration, cache, or log directories, platformdirs resolves the correct location for macOS, Windows, Linux/Unix, and Android.

## Features

<table>
<tr align=center><td align=left>

**Platform auto-detection**

Works on macOS, Windows, Linux, FreeBSD, OpenBSD, and Android – no configuration needed.

<td align=left>

**Convention-compliant**

XDG Base Directory Spec on Linux, ~/Library on macOS, AppData on Windows.

<tr align=center><td align=left>

**XDG variable support**

Honors XDG_DATA_HOME, XDG_CONFIG_HOME, and friends on both Linux and macOS.

<td align=left>

**str or Path**

Every directory has a _dir (str) and _path (Path) variant.

<tr align=center><td align=left>

**Auto-create directories**

Set ensure_exists=True to create directories on first access.

<td align=left>

**Version isolation**

Keep multiple app versions side by side with the version parameter.

</table>
