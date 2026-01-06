module gitoa.ru/go-4devs/console

go 1.24.0

require gitoa.ru/go-4devs/config v0.0.10

require (
	golang.org/x/mod v0.31.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/tools v0.40.0 // indirect
)

tool (
	gitoa.ru/go-4devs/config/cmd/config
	golang.org/x/tools/cmd/stringer
)
