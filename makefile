VERSION := 0.0.0
COMMIT  := $(shell git rev-parse --short HEAD)

VERSION_INJECT := -X github.com/sullrich84/preflight/app/build.Version=$(VERSION)
COMMIT_INJECT  := -X github.com/sullrich84/preflight/app/build.Commit=$(COMMIT)
LDFLAGS 	   := "$(VERSION_INJECT) $(COMMIT_INJECT)"

build:
	env GOOS=linux GOARCH=amd64 go build -tags netgo -ldflags $(LDFLAGS) -o dist/preflight-linux-$(VERSION) preflight.go
	env GOOS=darwin GOARCH=amd64 go build -tags netgo -ldflags $(LDFLAGS) -o dist/preflight-darwin-$(VERSION) preflight.go
	env GOOS=windows GOARCH=amd64 go build -tags netgo -ldflags $(LDFLAGS) -o dist/preflight-windows-$(VERSION).exe preflight.go