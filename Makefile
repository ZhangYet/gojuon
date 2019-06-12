.PHONY: all

all: server client

clean:
	rm bin/*

proto:
	protoc -I=./proto --go_out=plugins=grpc:rpc ./proto/*.proto

packr2:
	cd data && packr2 && cd -

server: proto packr2
	go build -ldflags "-X main.Version=`git describe --tags || echo "nightly"`" -o bin/gojuond cmd/gojuond/server.go

client: proto
	go build -ldflags "-X main.Version=`git describe --tags || echo "nightly"`" -o bin/gojuon cmd/gojuon/client.go

