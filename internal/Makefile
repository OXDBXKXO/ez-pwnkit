GO=$(shell which go)
GOCFLAGS = -buildmode=c-shared
GOLDFLAGS = -ldflags="-w -s"
TMPMAIN=tmp_main.go

.PHONY: PWN.so
PWN.so-go: payload.go
	sed 's/payload_template/main/' $< > $(TMPMAIN)
	$(GO) build -o PWN.so $(GOCFLAGS) $(GOLDFLAGS) $(TMPMAIN)
	rm $(TMPMAIN)
	rm PWN.h

PWN.so-c: payload.c
	gcc -fPIC -shared -g $< -o PWN.so

.PHONY: clean
clean:
	rm PWN.so