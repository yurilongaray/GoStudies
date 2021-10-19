# Overview

### Ellen Korbes
- https://www.youtube.com/watch?v=WiGU_ZB-u0w&t=237s&ab_channel=AprendaGo
- https://github.com/ellenkorbes/aprendago

### Cod3rCursos
https://www.udemy.com/course/curso-go/

## Docs
- https://gobyexample.com/
- https://golang.org/doc/effective_go


## Disappointments
- There is no Ternary and TryCatch in GoLang...

## Testing

```bash
	go test -v ./...
```
- "-v" = verbose
- "./..." = run in all folders and subfolders of your path

## Tools

- Go Doc (browser documentation from your code): godoc --http=:4000
- Go Fmt (format go files): gofmt -w ./...
- Go Vet (correctness → fetch for wrong code): go vet ./...
- Go Lint (suggestions → fetch for bad code style): golint ./...
- Go Test Flag (shows every kind of test): go help testflag
- Go Test Coverage:
	- go test --coverprofile coverage_report.out
	- go tool cover -html=coverage_report.out

- Design Pattern in Go: 
	- Generator: A simple function that creates generate go routines and returns a channel;
	- Multiplexer: Get two channels together into a single channel 

Obs.: When using "./..." the command will consider all folders and subfolders of your path