EXPLOIT=exploit

.PHONY: default
default: build

.PHONY: malicious_lib
malicious_lib:
	make -C ./internal PWN.so

build: malicious_lib
	go build -o $(EXPLOIT) ./cmd/main.go

.PHONY: clean
clean:
	make -C ./internal clean
	rm $(EXPLOIT)