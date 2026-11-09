module github.com/go-4devs/console

go 1.27.3

require github.com/go-4devs/config v0.0.17

require (
	gitoa.ru/go-4devs/config v0.0.11 // indirect
	golang.org/x/mod v0.42.0 // indirect
	golang.org/x/sync v0.25.0 // indirect
	golang.org/x/tools v0.52.0 // indirect
)

tool (
	github.com/go-4devs/config/cmd/config
	golang.org/x/tools/cmd/stringer
)
