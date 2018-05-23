## Introduction

This program shows contents of a plain text file line by line in a certain period. It's designed for people who is learning new languages or anyone who need to remember or review a lot of unrelated small things.

The program is pretty simple to use, and this is how it looks:
![software appearence](./doc/eg.png)

Start the software, choose the file to display, then you can put it into the corner of your screen, and start your daily work.

![The running software](./doc/eg.gif)

If you want to change the frequency(defaults to 2 seconds for 1 line), just write down the time you want(have to be an interger) and click "Change" button.

This program support Windows and Linux. It's tested on Linux4.4.0 x86_64 and Windows7 x64.


## Build

First you need to have [Go language][1] tools installed. For Chinese users, you can get it from [here][2]. After that, enter the program directory, build it with this:

```sh
go build
```

For Windows, to eliminate the useless cmd window, you need to add some build flags:
```sh
go build -ldflags -H=windowsgui
```

## Prebuild

There is a prebuild binary for MS Windows. [Click][3] to download it.


[1]: https://golang.org/
[2]: https://golang.google.cn/
[3]: https://github.com/wallacegibbon/lineslide/blob/master/prebuild/lineslide.exe.bz2?raw=true

