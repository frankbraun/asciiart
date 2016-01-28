all:
	go install -v github.com/frankbraun/asciiart/...

templates/exampleart.svg: templates/exampleart.tmpl
	aa2svg -i $< -o $@ -f

.PHONY: generate test
generate: all templates/*.tmpl templates/exampleart.svg
	aatmpl

test: all
	go test -v github.com/frankbraun/asciiart/aa2d
