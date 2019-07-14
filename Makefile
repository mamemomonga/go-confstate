NAME := confstate-example

SRCS := $(shell find . -type f -name '*.go')

bin/$(NAME): $(SRCS)
	go build -o $@ ./examples

clean:
	rm -rf bin
	rm -f config.yaml state.json

