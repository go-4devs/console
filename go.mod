module github.com/go-4devs/console

go 1.27.3

require github.com/go-4devs/config v0.1.0

require (
	github.com/go-4devs/config/cmd/config v0.0.0-20261109185843-96a65123fd7f // indirect
	golang.org/x/mod v0.42.0 // indirect
	golang.org/x/sync v0.25.0 // indirect
	golang.org/x/tools v0.52.0 // indirect
)

tool (
	github.com/go-4devs/config/cmd/config
	golang.org/x/tools/cmd/stringer
)
