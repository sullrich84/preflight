module github.com/sullrich84/preflight

go 1.15

require (
	github.com/ahmetalpbalkan/go-cursor v0.0.0-20131010032410-8136607ea412
	github.com/gosuri/uilive v0.0.4
	github.com/logrusorgru/aurora/v3 v3.0.0
	github.com/spf13/cobra v1.1.1
	github.com/thoas/go-funk v0.7.0
)

replace github.com/sullrich84/preflight/app => ./app

replace github.com/sullrich84/preflight/cmd => ./cmd

replace github.com/sullrich84/preflight/preflight => ./preflight

replace github.com/sullrich84/preflight/terminal => ./terminal
