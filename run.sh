cp ./lib/darwin/libledLib.dylib /usr/local/lib/
GOPATH=$GOPATH:./
go run main.go
