.PHONY: build clean test

build: 
	go build

clean:
	rm -f gosp *~

test: build
	@./test.sh
