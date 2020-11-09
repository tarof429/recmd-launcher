BASEPATH = $(PWD)
GOPATH = $(PWD)/build
DISTPATH = $(PWD)/dist
RECMD_DMN_GIT_REPO = https://github.com/tarof429/recmd-dmn.git
RECMD_CLI_GIT_REPO = https://github.com/tarof429/recmd-cli.git

default: all

.PHONY: recmd-dmn recmd-cli clean package repackage

all: recmd-dmn recmd-cli repackage package

recmd-dmn:
	@mkdir -p $(GOPATH)
	@cd $(GOPATH); [ -d recmd-dmn ] || git clone $(RECMD_DMN_GIT_REPO)
	@cd $(GOPATH)/recmd-dmn && git pull
	@cd $(GOPATH)/recmd-dmn && make clean build test || exit 1

recmd-cli:
	@mkdir -p $(GOPATH)
	@cd $(GOPATH); [ -d recmd-cli ] || git clone $(RECMD_CLI_GIT_REPO)
	@cd $(GOPATH)/recmd-cli && git pull
	@cd $(GOPATH)/recmd-cli && make clean build test || exit 1

package:
	@mkdir -p $(DISTPATH)/bin
	@mkdir -p $(DISTPATH)/logs
	@mkdir -p $(DISTPATH)/conf
	@cp $(GOPATH)/recmd-dmn/recmd-dmn $(DISTPATH)/bin
	@cp $(GOPATH)/recmd-cli/recmd-cli $(DISTPATH)/bin
	@cd $(DISTPATH); ln -sf bin/recmd-cli recmd
	@cd $(DISTPATH); zip --symlinks -r $(BASEPATH)/recmd-launcher.zip .

repackage:
	go build -o dist/bin/repackage repackage.go 
	
clean:
	@if [ -d build ]; then chmod -R 777 build; fi
	@rm -rf build dist recmd-launcher.zip