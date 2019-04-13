.PHONY: build clean test install

build: 
	go build

install: 
	go install github.com/Matts966/gosp/cmd/gosp

clean:
	rm -f gosp *~

test: install
	@./test.sh
