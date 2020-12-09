build:
	env GOOS=linux GOARCH=amd64 go build -o dist/preflight-linux preflight.go
	env GOOS=darwin GOARCH=amd64 go build -o dist/preflight-macos preflight.go