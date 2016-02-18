#__go=$(shell which go)
__go=/Users/a12057/homebrew/Cellar/go/1.5.2/bin/go
__vendor_dir=vendor
__gopath=$(shell pwd)/$(__vendor_dir)

all: deps build

deps:
	@mkdir -p $(__gopath)
	GOPATH=$(__gopath) $(__go) get google.golang.org/api/vision/v1
	GOPATH=$(__gopath) $(__go) get google.golang.org/api/googleapi/transport

run:
	GOPATH=$(__gopath) $(__go) run main.go

build:
	GOPATH=$(__gopath) $(__go) build -o google-vision main.go
