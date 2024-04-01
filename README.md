# Golang

GOPATH
GOROOT
GOBIN
GOOS
GOARCH

# 25 keywords

break,case,chan,continue,const,default,defer,else,fallthrough,for,func,go,goto,if,import,interface,map,package,range,return,select,struct,switch,type,var


# builtin 

complex,println, print, cap, len, copy, append, delete, make, new , panic, recover,close


# Go mod

```
go mod init demo
```

# build

```
go build main.go # normal build
go build -o hello main.go # build with a specific output file
go build -ldflags "-s -w" -o hello1 main.go # stripe down build.
```

# install
- it installs binaries from your local system or from the repos into the GOBIN directory

```
go install .
go install github.com/JitenPalaparthi/mathshapes@latest
```
# cross compilation

```
go tool dist list
```
GOOS/GOARCH
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
freebsd/arm64
freebsd/riscv64
illumos/amd64
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/loong64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
openbsd/ppc64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
wasip1/wasm
windows/386
windows/amd64
windows/arm
windows/arm64

## cross compilation

```
GOOS=windows GOARCH=arm64 && go build -ldflags="-s -w" -o hello-windows.exe main.go
```

# Go Escape analysis

```
go build -o app main.go -gcflags="-m"
```

# Test go code

```go test -timeout 30s -run ^TestValidateFail$ demo/models```

# All the files to be tested..

```
go test ./...
```
# Test a specific file

```
go test models/person_test.go 
```
# To Run file test

```
go test -timeout 30s -run ^(TestValidateSuccess|TestValidateFail)$ demo/models
```


# To write test coverage

```
go test -timeout 30s  demo/models -coverprofile cover.out
```

# To write to html outout of test coverate

```
go tool cover -html=cover.out 
```
# To check test coverage per func

```
go tool cover -func=cover.out
```


