GO=$(shell which go)
GOCFLAGS = -buildmode=c-shared
GOLDFLAGS = -ldflags="-w -s"
TMPMAIN=tmp_main.go

.PHONY: PWN.so
PWN.so: payload.go
	sed 's/payload_template/main/' $< > $(TMPMAIN)
	$(GO) build -o PWN.so $(GOCFLAGS) $(GOLDFLAGS) $(TMPMAIN)
	rm $(TMPMAIN)
	rm PWN.h

.PHONY: clean
clean:
	rm PWN.so