all: clean
	go build

.PHONY: clean
clean:
	rm -f ./server
