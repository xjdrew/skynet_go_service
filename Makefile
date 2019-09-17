all:
	CGO_CFLAGS='-I /home/tiger/work/tadpole/template/build/include' CGO_LDFLAGS='-Wl,--unresolved-symbols=ignore-all' go build -buildmode=c-shared -o gos.so gos.go
clean:
	-rm gos.go
