build:
	$${HOME}/go/bin/packr build

run:
	./knowhow search

clean:
	rm debug.db knowhow

all: build run

