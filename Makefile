.PHONY: build clean test

build: 
	go build -o minilisp minilisp.go

clean:
	rm -f minilisp *~

test: minilisp
	@./test.sh
