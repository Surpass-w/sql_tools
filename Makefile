export GOPROXY=https://goproxy.cn,direct

LDFLAGS="-w -s"

all: clear build

clear:
	rm -rf bin/

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags=${LDFLAGS} -o bin/linux/sql_tools main.go
#	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags=${LDFLAGS} -o bin/darwin/tools main.go
#	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags=${LDFLAGS} -o bin/darwin/tools main.go