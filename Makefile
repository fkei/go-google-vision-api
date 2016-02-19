__go=$(shell which go)
__vendor_dir=vendor
__gopath=$(shell pwd)/$(__vendor_dir)

API_KEY=$(api-key)

all: deps build

deps:
	@mkdir -p $(__gopath)
	GOPATH=$(__gopath) $(__go) get google.golang.org/api/vision/v1
	GOPATH=$(__gopath) $(__go) get google.golang.org/api/googleapi/transport

run:
	GOPATH=$(__gopath) $(__go) run main.go -api-key=$(API_KEY) -path=test.png

build:
	GOPATH=$(__gopath) $(__go) build -o google-vision main.go
