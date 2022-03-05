# median-cli
A CLI which returns the median of given numbers

#### Legend
- `<name>` Required
- `[name]` Optional
- `<names[]>` Required array
- `[names[]]` Optional array
- `./program[.exe]` Optional file extension (based on OS)


## Build
1. `git clone https://github.com/thestevexyz/median-cli.git` or download repository
2. Change directory to repository
2. `go build -ldflags "-s -w" -trimpath .`
3. Run CLI with `./median-cli[.exe] <numbers[]>`

## Testing
`go test`

## OS/ARCH Support
Only Linux (*Ubuntu 20.04*, *Debian 10*) and *Windows 10* are officially supported.
However, other builds are also possible, but not officially supported.

A bash script `build-executables.bash` is provided to build for multiple OS's and Architectures.

Usage: `./build-executeables.bash main.go median-cli`