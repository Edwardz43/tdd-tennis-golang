# Tennis for TDD

Golang version

master branch includes only project initialization

all practice history included in practice branch

you can use `make test` to show test coverage detail

or use `make unnittest` to show simple test result if Makefile supported by your environment

otherwise, directly use commands:
```
go test -v -cover -covermode=atomic ./...
```
or
```
go test -short  ./...
```