# median-cli
A CLI which returns the median of given numbers

#### Legend
- `<name>` Required
- `[name]` Optional
- `<names[]>` Required array
- `[names[]]` Optional array
- `./program[.exe]` Optional file extension (based on OS)


## Build
1. `git pull https://github.com/thestevexyz/median-cli.git` or download repository
2. Change directory to repository
2. `go build -ldflags "-s -w" -trimpath .`
3. Run CLI with `./median-cli[.exe] <numbers[]>`

## Testing
`go test`

