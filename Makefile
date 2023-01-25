runserver:
	./dev_run.sh

tsgen: 
	npm run build

tstypegen:
	tygo generate

dev_setup:
	go install github.com/tkrajina/typescriptify-golang-structs/tscriptify
	go install github.com/cespare/reflex@latest

install:
	go get ./...
	go mod vendor
