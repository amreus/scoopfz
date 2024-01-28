SHELL = cmd

CFLAGS = -Wall -Wextra
CC = gcc

.SUFFIXES:


gen-list.exe: gen-list.c
	$(CC) $(CFLAGS) -o gen-list.exe gen-list.c C:\Users\Jim\repos\cJSON\b\libcjson.a

t: gen-list.exe
	gen-list.exe
	head app-list
	wc -l app-list

clean:
	del gen-list.exe
	del app-list-*
