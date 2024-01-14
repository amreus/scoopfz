SHELL = cmd

scoopfz.exe : scoopfz.go
	go build

.PHONY : clean fmt

fmt:
	gofmt -w scoopfz.go

clean :
	del scoopfz.exe
