# Tutorial

In this tutorial, we will set up platformdirs for a small application that stores user preferences, caches API responses, and writes log files – all to the correct platform-specific directories.

## Installation

Install `platformdirs` from go.jcbhmr.com:

```sh
go get go.jcbhmr.com/platformdirs
```

Verify the installation:

```go
import "go.jcbhmr.com/platformdirs"
fmt.Println(platformdirs.Version)
```

## Finding your first directory


