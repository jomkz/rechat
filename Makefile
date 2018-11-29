.PHONY: build clean image run

build:
	mkdir -p ./build
	GOARCH=amd64 GOOS=linux go build -o ./build/rechat github.com/jmckind/rechat/cmd

clean:
	rm -fr ./build

image:
	docker build -t jmckind/rechat:latest .

run:
	./build/rechat
