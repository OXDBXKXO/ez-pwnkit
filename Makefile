EXPLOIT=exploit

default: build_c

.PHONY: go_malicious_lib
go_malicious_lib:
	make -C ./internal PWN.so-go

.PHONY: c_malicious_lib
c_malicious_lib:
	make -C ./internal PWN.so-c

.PHONY: build_go
build_go: go_malicious_lib
	go build -o $(EXPLOIT) ./cmd/main.go

.PHONY: build_c
build_c: c_malicious_lib
	go build -o $(EXPLOIT) ./cmd/main.go

.PHONY: clean
clean:
	make -C ./internal clean
	rm $(EXPLOIT)