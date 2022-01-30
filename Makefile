EXPLOIT=exploit

.PHONY: default
default: build

exploit:
	make -C ./internal PWN.so

build: exploit
	go build -o $(EXPLOIT) ./cmd/main.go

.PHONY: clean
clean:
	make -C ./internal clean
	rm $(EXPLOIT)