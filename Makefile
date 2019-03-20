.PHONY: build clean test

build: 
	go build -o gosp gosp.go

clean:
	rm -f gosp *~

test: build
	@./test.sh
