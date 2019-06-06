.PHONY: all

all:
	go build -ldflags "-X main.Version=`git describe --tags || echo "nightly"`" -o bin/gojuon gojuon.go
