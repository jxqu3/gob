# gob: go cross-platform build tool
What is `gob`? `gob` is a small (<2MB) cli tool to make it easier to cross-build in go. Instead of having to set your env variables every time, specially annoying in windows, you just do `gob --os=linux` or the target OS you want (as long as it is supported by the golang compiler)

### Features
1. Easy cross compile without env variables
2. List available OSs and architectures
3. Shorter than `go build .`, just run `gob`
4. Only 2MB on windows!

### Install:
1. Run `go install github.com/checkm4ted/gob/v2@master`

### Usage:
`gob --help`  
```cs
Commands:
	--help: show this help
    --os: set target platform OS
    --arch: set target platform CPU architecture
    --out: set output file
    --cgo: enable cgo
    --light: sets -w and -s ldflags
    --ldflags: set linker flags
    --env: set other env vars
    --list-os: list available OS
    --list-arch: list available CPU architectures
    Example:
    gob --os=linux --arch=amd64 --out=linux_amd64 main.go
```

Note: use --light AFTER --ldflags, else it won't be affected.

If you just wanna build your project to the default platform, you can just do this:  
`gob`  
This  will execute `go build .` under the hood. 

To build to a different platform in go, you have to set `GOOS` and `GOARCH` env variables first, which is somewhat easy in linux, but in windows it's really annoying. With gob you just need to pass the arguments  
`gob --os=linux`. If you don't specify the arch or OS it will use the current env values you have already set.


### Why?
To learn to make a CLI tool and save some minutes of my time in the future

### Requirements:
You need to have go installed as this tool is ***NOT A COMPILER***, it calls go build, this just facilitates the cross platform options
