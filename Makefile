all:
	go install -v github.com/frankbraun/asciiart/cmd/...

.PHONY: test
test:
	go test -v github.com/frankbraun/asciiart/aa2d
