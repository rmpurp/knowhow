build:
	$${HOME}/go/bin/packr build  --tags "fts5"

run:
	./knowhow search

clean:
	rm -f debug.db knowhow

all: clean build run

