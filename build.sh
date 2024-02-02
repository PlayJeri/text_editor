# Build Mac & Windows binaries
export GOOS=windows
go build -o ISO.exe src/main.go
export GOOS=darwin
go build -o ISO src/main.go