SHELL = cmd

CFLAGS = -Wall -Wextra -g
CC = gcc

.SUFFIXES:


gen-list.exe: gen-list.c
	$(CC) $(CFLAGS) -o gen-list.exe gen-list.c ..\..\..\cJSON\b\libcjson.a


.PHONY : t clean

t: gen-list.exe
	gen-list.exe

clean:
	del gen-list.exe
	del app-list-*
