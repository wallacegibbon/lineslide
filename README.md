## Introduction

This program is based on the [andlabs/ui][1] project(A Platform-native GUI library for Go).


## Building

To eliminate the useless cmd window, you need to add some build flags:
```sh
go build -ldflags -H=windowsgui
```

## Prebuild

There is a prebuild binary for MS Windows. [Click][2] to download it.


[1]: https://github.com/andlabs/ui/
[2]: https://github.com/wallacegibbon/lineslide/blob/master/prebuild/lineslide.exe.bz2?raw=true
