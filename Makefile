build_agent:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o survive_linux -ldflags "-s -w" cmd/client/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o survive.exe -ldflags "-s -w -H windowsgui" cmd/client/main.go

build_server:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o survive_server_linux -ldflags "-s -w" cmd/server/main.go