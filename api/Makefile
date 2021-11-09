all: zapi_unix.go

zapi_unix.go: api.go
	./mksyscall_unix.pl $^ \
		| gofmt \
		> $@
