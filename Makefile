ifeq ($(OS),Windows_NT)
	SHELL = cmd
endif

scoopfz : scoopfz.go
	go build $^

.PHONY : clean fmt

fmt:
	gofmt -w scoopfz.go

clean :
	del scoopfz.exe
