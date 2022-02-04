## Number Split Randomly

The sample code written in `Go lang` 

### Executables

This repo already have executables for Linux and windows, after clone the repo just navigate to the source the follow the steps based on your OS.

For Linux from CLI

`./randomsplit`

For windows run the `randomsplit.exe`

### Run & Build

Run the application from CLI ([Go](https://go.dev/doc/install) must be installed on the local machine or use [docker container for Go](https://github.com/Jobinjose01/Go-MongoDB-Docker))

`go run main.go`

##### Build Commands

For native OS build

`go build`

For windows build from Linux 
```
 GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build
```

### Test Cases

For running test cases (test cases can be modified in the main_test.go file as per the requirements)

`go test`

or 

`go test -v`

### Coverage Test

`go test -coverprofile=coverage.out`

HTML view

`go tool cover -html=coverage.out`

### Benchmarks

`go test -bench=.`

### Access The CLI with Docker

`docker ps`

`docker exec -it <go container ID> /bin/sh` 