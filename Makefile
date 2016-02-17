all:
	go install -v github.com/frankbraun/asciiart
	go install -v github.com/frankbraun/asciiart/cmd/...
	go install -v github.com/frankbraun/asciiart/svg/...

templates/%.svg: templates/%.tmpl
	aa2svg -i $< -o $@ -f

.PHONY: generate test
generate: templates/*.tmpl templates/exampleart.svg
	go generate -v github.com/frankbraun/asciiart/svg/path
	goimports -w svg/path/grammar.go
	go install -v github.com/frankbraun/asciiart/util/cmd/aatmpl
	aatmpl

test: all
	gocheck -g -c -e talk
