module gitoa.ru/go-4devs/console/example

go 1.22

toolchain go1.24.1

replace gitoa.ru/go-4devs/console => ../

replace gitoa.ru/go-4devs/console/input/cfg => ../input/cfg

require (
	gitoa.ru/go-4devs/config v0.0.0-20210427173104-3ba6b4c71578
	gitoa.ru/go-4devs/console v0.1.0
	gitoa.ru/go-4devs/console/input/cfg v0.0.1
)
