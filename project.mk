# Copyright 2018-2019 RICHARD Didier
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Adapted from https://github.com/vincentbernat/hellogopher/blob/master/Makefile
#
# To test without executing commands :
# $ make -n
# To debug :
# $ make --debug=[a|b|v|i|j|m]
#

# Print a * in blue :
#S        = $(shell printf "\033[34;1m*\033[0m")
# Print a * :
S        = $(shell printf "*")
# Print a TAB :
#T        = $(shell printf "\t")
# Print two spaces :
T        = $(shell printf "  ")
# Print a BACKSPACE :
#B        = $(shell printf "\b")
# Reverse a list (thx to https://stackoverflow.com/questions/52674/simplest-way-to-reverse-the-order-of-strings-in-a-make-variable) :
ReverseWordList = $(if $(wordlist 2,2,$(1)),$(call ReverseWordList,$(wordlist 2,$(words $(1)),$(1))) $(firstword $(1)),$(1))

DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell [ -f $(PKGDIR)/version.go ] && cat $(PKGDIR)/version.go | grep "\w Version" | sed -e 's/[^"]*"\([^"]*\)"/\1/' || \
                git describe --tags --always --dirty --match=v* 2> /dev/null || \
                echo "0.0.0")
# one can override GOPATH by passing it to the command line. This is usefull
# when installing this package for production : make GOPATH=${GOPATH}
# If not defined set to current working directory
GOPATH  ?= .
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/src/$(PACKAGE)
PROJ_PATH=.

GO       = go
GODOC    = go doc
DEP      = dep
GOLINT   = golint
GOI18N   = $(BIN)/goi18n

# in case of cross-compiling :
CGO_ENABLED     = 1
PKG_CONFIG_PATH = $(PROJ_PATH)/usr/local/lib/pkgconfig
LDFLAGS         =-ldflags '-X $(PACKAGE).Version=$(VERSION) -X $(PACKAGE).BuildDate=$(DATE)'
PROJ_LIB_PATH   = $(PROJ_PATH)/usr/local/lib
LD_LIBRARY_PATH =$(PROJ_LIB_PATH)
PROJ_LIB        = $(PROJ_PATH)/usr/local/share/proj

export GOPATH
export CGO_ENABLED
export PKG_CONFIG_PATH
export LD_LIBRARY_PATH
export PROJ_LIB

pkgPath=$(PKGROOT)
# remember : one can pass PACKAGE on the command-line like $ make PACKAGE=${PKGROOT}/subpkg
pkgName=$(shell basename $(PACKAGE))
ifneq ($(shell basename $(pkgPath)),$(pkgName))
pkgPfx=.$(PACKAGE:$(pkgPath)%=%)/
else
pkgPfx=./
endif
pkgFile=$(pkgPfx)$(pkgName)
coverOut=$(pkgFile)-coverprofile.out
coverHtml=$(pkgFile)-cover.html
pkgDoc=$(pkgFile)-doc.txt

# Default target (first target not starting with .)
main:: ; $(info $(S) Building …)
ifdef TRACE
main:: check-env print-env tools vendor packages test install
else
main:: check-env tools vendor packages test install
endif

# use make TRACE=1 to execute print-env target
ifdef TRACE
print-env:
	$(info $(T)PATH=$(PATH))
	$(info $(T)GOPATH=$(GOPATH))
	$(info $(T)CGO_ENABLED=$(CGO_ENABLED))
	$(info $(T)PKG_CONFIG_PATH=$(PKG_CONFIG_PATH))
	$(info $(T)LD_LIBRARY_PATH=$(LD_LIBRARY_PATH))
	$(info $(T)PROJ_LIB=$(PROJ_LIB))
	$(info $(T)pkgPath=$(pkgPath))
	$(info $(T)pkgName=$(pkgName))
	$(info $(T)pkgPfx=$(pkgPfx))
	$(info $(T)pkgFile=$(pkgFile))
	$(info $(T)coverOut=$(coverOut))
	$(info $(T)coverHtml=$(coverHtml))
	$(info $(T)pkgDoc=$(pkgDoc))
endif

# Source directory
$(BASE):: check-env

# Check environment
# go is somewhere in the user's PATH :
check-env:: ; $(info $(T)$(T)$(S) checking for $(GO) …)
ifeq ($(shell command -v "$(GO)" >/dev/null 2>&1 || printf "0"),0)
check-env::
	$(error $(GO) is not in $$PATH.  Visit https://golang.org/dl/)
else
GOOS    ?= $(shell uname -s | tr A-Z a-z)
GOARCH  ?= $(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m)))
check-env::
	$(info $(T)$(T)… found with GOARCH="$(GOARCH)" and GOOS="$(GOOS)")
endif

$(BIN):
	@mkdir -p $@
# the order-only pre-requisite $(BASE) is executed first if needed
$(BIN)/%:: $(BIN) | $(BASE) ; $(info $(T)$(T)$(S) installing $(@F) from $(REPOSITORY) …)
	@ tmp=$$(mktemp -d) ; \
		(env GOPATH=$$tmp $(GO) get -u $(REPOSITORY) && cp $$tmp/bin/* $(BIN)/.) || ret=$$?; \
		rm -rf $$tmp ; exit $$ret
$(BIN)/%:: ; $(info $(T)$(T) $(@F) installed !)

# Tools :
tools: tools-prolog dep golint goi18n tools-epilog
tools-prolog: $(BASE)
	$(info $(T)$(S) Checking/Installing tools …)
tools-epilog:
	$(info $(T)… done)

dep:: check-env ; $(info $(T)$(T)$(S) checking for $(DEP) …)
ifeq ($(shell command -v "$(DEP)" >/dev/null 2>&1 || printf "0"),0)
dep:: $(BIN)/dep
$(BIN)/dep: REPOSITORY=github.com/golang/dep/cmd/dep
else
dep::
	$(info $(T)$(T)… found)
endif

golint:: check-env ; $(info $(T)$(T)$(S) checking for $(GOLINT) …)
ifeq ($(shell command -v "$(GOLINT)" >/dev/null 2>&1 || printf "0"),0)
golint:: $(BIN)/golint
$(BIN)/golint: REPOSITORY=github.com/golang/lint/golint
else
golint::
	$(info $(T)$(T)… found)
endif

goi18n:: check-env ; $(info $(T)$(T)$(S) checking for $(GOI18N) …)
ifeq ($(shell command -v "$(GOI18N)" >/dev/null 2>&1 || printf "0"),0)
goi18n:: $(BIN)/goi18n
$(BIN)/goi18n: REPOSITORY=github.com/nicksnyder/go-i18n/goi18n
else
goi18n::
	$(info $(T)$(T)… found)
endif
goi18n::

# Package test
ifndef WITHOUTPKGS
test::
	@for p in $(SUBPKGS) ; do \
		$(MAKE) GOPATH=$(GOPATH) -f ./$${p}/Makefile test ;\
	done
endif
test:: test-prolog tools vendor build-test test-epilog

test-prolog: $(BASE)
	$(info $(S) Building environment and testing …)

test-epilog:
	$(info … tested)

ifndef WITHOUTPKGS
clean::
	@for p in $(SUBPKGS) ; do \
		$(MAKE) GOPATH=$(GOPATH) -f ./$${p}/Makefile clean ;\
	done
endif
clean:: clean-prolog clean-tmp clean-i18n clean-package clean-epilog

clean-prolog:
	$(info $(T)$(S) Cleaning …)

clean-tmp:
	@rm -f "$(coverOut)" "$(coverHtml)" "$(pkgDoc)"

clean-i18n:
	@rm -f $(pkgPfx)i18n/*.all.json $(pkgPfx)i18n/*.untranslated.json

clean-package:
	@rm -f Gopkg.awk
	@$(GO) clean $(PACKAGE) ; true # the PACKAGE may have not been yet installed ...

clean-epilog:
	$(info $(T)… cleaned)

build-test: build-i18n test-lint

ifneq ("$(wildcard $(pkgPfx)i18n/*.toml)","")
build-i18n: build-i18n-prolog goi18n $(pkgPfx)i18n/en-us.all.json $(pkgPfx)i18n/fr-fr.all.json build-i18n-epilog

build-i18n-prolog: $(BASE)
	$(info $(T)$(S) Generating i18n files …)

build-i18n-epilog:
	$(info $(T)… generation done)

%.all.json: %.toml
	$(eval lang=$(shell basename `echo $*`))
	$(eval i18n=$(shell dirname `echo $@`))
	@$(GOI18N) -sourceLanguage $(lang) -outdir $(i18n) ./$(i18n)/$(lang).toml
else
build-i18n:
	$(info $(T)$(S) Nothing to do for i18n)
endif

test-lint: test-lint-prolog count cover test-lint-epilog

test-lint-prolog:
	$(info $(T)$(S) Preparing, Calculating coverage and Linting …)

count:
	@env CGO_ENABLED=$(CGO_ENABLED) PKG_CONFIG_PATH=$(PKG_CONFIG_PATH) LD_LIBRARY_PATH=$(LD_LIBRARY_PATH) PROJ_LIB=$(PROJ_LIB) $(GO) test $(GOTESTOPTS) \
		$(LDFLAGS) \
		-covermode=count -coverprofile="$(coverOut)" $(PACKAGE)

cover:
	@$(GO) tool cover -func="$(coverOut)"
	@$(GO) tool cover -html="$(coverOut)" -o "$(coverHtml)"
	@rm -f "$(coverOut)"
	@echo firefox "$(coverHtml)" to see $(PACKAGE) code\'s coverage \(heat map\)
	@$(GOLINT) $(PACKAGE)
	@$(GO) vet $(PACKAGE)

test-lint-epilog:
	$(info $(T)$(S) … tests, coverage and lint done)

doc: doc-prolog doc-godoc doc-epilog

doc-prolog:
	$(info $(T)$(S) Generating text document in $(pkgDoc) …)

doc-godoc:
	@$(GODOC) $(PACKAGE) | sed '1d' > "$(pkgDoc)"

doc-epilog:
	$(info $(T)… documentation done)

# Installation of the package
ifndef WITHOUTPKGS
install::
	@for p in $(SUBPKGS) ; do \
		$(MAKE) GOPATH=$(GOPATH) -f ./$${p}/Makefile install ;\
	done
endif
install:: install-prolog clean build-i18n install-package doc install-epilog

install-prolog: $(BASE)
	$(info $(S) Installing $(PACKAGE) …)

install-package:
	@env GOARCH=$(GOARCH) GOOS=$(GOOS) \
		CGO_ENABLED=$(CGO_ENABLED) PKG_CONFIG_PATH=$(PKG_CONFIG_PATH) LD_LIBRARY_PATH=$(LD_LIBRARY_PATH) PROJ_LIB=$(PROJ_LIB) $(GO) install \
			$(LDFLAGS) $(PACKAGE)

install-epilog:
	$(info … installation done)

# always execute these targets when invoked
.PHONY: all main tools test clean build-test doc install help

