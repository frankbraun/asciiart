all:
	go install -v github.com/frankbraun/asciiart/...

.PHONY: generate test
generate: all templates/*.tmpl
	aatmpl

test: all
	go test -v github.com/frankbraun/asciiart/aa2d
