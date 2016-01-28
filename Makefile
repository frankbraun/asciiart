all:
	go install -v github.com/frankbraun/asciiart/...

templates/exampleart.svg: templates/exampleart.tmpl
	aa2svg -i $< -o $@ -f

.PHONY: generate test
generate: templates/*.tmpl templates/exampleart.svg
	go install -v github.com/frankbraun/asciiart/util/cmd/aatmpl
	aatmpl

test: all
	go test -v github.com/frankbraun/asciiart/aa2d
