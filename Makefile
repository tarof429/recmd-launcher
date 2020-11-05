BASEPATH = $(PWD)
GOPATH = $(PWD)/build
DISTPATH = $(PWD)/dist
RECMD_DMN_GIT_REPO = https://github.com/tarof429/recmd-dmn.git
RECMD_CLI_GIT_REPO = https://github.com/tarof429/recmd-cli.git

default: all

.PHONY: recmd-dmn recmd-cli clean package

all: recmd-dmn recmd-cli package

recmd-dmn:
	@mkdir -p $(GOPATH)
	@cd $(GOPATH); [ -d recmd-dmn ] || git clone $(RECMD_DMN_GIT_REPO)
	@cd $(GOPATH)/recmd-dmn && git pull
	@cd $(GOPATH)/recmd-dmn && make clean build test install || exit 1

recmd-cli:
	@mkdir -p $(GOPATH)
	@cd $(GOPATH); [ -d recmd-cli ] || git clone $(RECMD_CLI_GIT_REPO)
	@cd $(GOPATH)/recmd-cli && git pull
	@cd $(GOPATH)/recmd-cli && make clean build test install || exit 1

package:
	@mkdir -p $(DISTPATH)
	@mkdir -p $(DISTPATH)/bin
	@mkdir -p $(DISTPATH)/logs
	@touch $(DISTPATH)/logs/.ignore
	@cd $(DISTPATH); ln -sf bin/recmd-cli recmd
	@cp $(GOPATH)/bin/* $(DISTPATH)/bin
	@cd $(DISTPATH); zip --symlinks -r -D $(BASEPATH)/recmd-launcher.zip .

clean:
	@if [ -d build ]; then chmod -R 777 build; fi
	@rm -rf build dist recmd-launcher.zip