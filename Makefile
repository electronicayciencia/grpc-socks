version=$(shell git rev-parse --short HEAD)
buildAt=$(shell date "+%Y-%m-%d %H:%M:%S %Z")


build:
	rm -rf exec_bin
	mkdir exec_bin
	#GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=$(version) -X 'main.buildAt=$(buildAt)'" -o ./exec_bin/server-linux ./server/*.go

	GOOS=linux GOARCH=amd64 go build -o ./exec_bin/client-linux ./client/*.go

	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./exec_bin/server-linux ./server/*.go
	GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ./exec_bin/server-arm ./server/*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./exec_bin/server-alpine ./server/*.go


idl:
	rm -rf pb/*.pb.go
	protoc -I=. pb/*.proto --go_out=plugins=grpc:.

