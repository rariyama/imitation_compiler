#CFLAGS=-std=c11 -g -static

build:
	go build main.go

test: build
		./test.sh


clean:
		rm -f main *.o *~ tmp*

.PHONY: test clean
