# Copyright 2017-2019 Didier Richard
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

# top package
PKGROOT := osgeo.org/proj
PKGNAME  =
PKGDIR   = ./$(PKGNAME)
# one can invoke with a sub-package by passing PACKAGE to the command line
PACKAGE ?= $(PKGROOT)
SUBPKGS  =

include project.mk

# Default target (first target not starting with .)
all: main

ifndef WITHOUTPKGS
packages:
	@for p in $(SUBPKGS) ; do \
		$(MAKE) GOPATH=$(GOPATH) -f ./$${p}/Makefile ;\
	done
else
packages:
endif

main:: ; $(info … built done !)

ifeq ($(PACKAGE),$(PKGROOT))
# Dependencies through dep :
dep-init: $(BASE) ; $(info $(T)$(T)$(S) creating dependencies from scratch …)
	@cd $(BASE) && ( [ ! -f Gopkg.toml ] && $(DEP) init || touch Gopkg.toml )
dep-install: Gopkg.toml ; $(info $(T)$(T)$(S) retrieving dependencies …)
	@$(DEP) ensure

# protect lf -> \n\
# protect "  -> \"
# protect $  -> $$
# protect $0 from being replaced by /bin/sh -> \$$0
# protect \n -> \\\n
define AWKSCRIPT
#!/usr/bin/awk -f\n\
BEGIN{\n\
  P=0\n\
  U=\"\"\n\
}\n\
/^#/{\n\
  next\n\
}\n\
/^$$/{\n\
  next\n\
}\n\
/^\[\[projects\]\]/{\n\
  P=1\n\
  next\n\
}\n\
{\n\
  if (P==1 && match(\$$0,/\s*name\s*=\s*\"([^\"]*)\"/,v)>0) {\n\
    U=v[1]\n\
    next\n\
  }\n\
  if (P==1 && U!=\"\" && match(\$$0,/\s*packages\s*=\s*\[([^]]*)\]/,p)>0) {\n\
    n = split(p[1],pks,/,/)\n\
    if (n>0) {\n\
      pkg=pks[1]\n\
      gsub(/\"/,\"\",pkg)\n\
      if (pkg == \".\") {\n\
        printf \"%s\\\n\", U\n\
      } else {\n\
        printf \"%s/%s\\\n\", U, pkg\n\
      }\n\
      P=0\n\
      U=\"\"\n\
      next\n\
    }\n\
  }\n\
}\n\
END{}\n
endef

Gopkg.awk: Gopkg.toml Gopkg.lock | $(BASE) $(DEP)
	@echo "$(AWKSCRIPT)" > $@

vendor: Gopkg.toml Gopkg.lock Gopkg.awk | $(BASE) $(DEP) ; $(info $(T)$(T)$(S) installing dependencies …)
	@$(DEP) ensure -update
	@for p in `awk -f Gopkg.awk Gopkg.lock` ; do [ -d vendor/$$p ] && cd vendor/$$p && $(GO) install . && cd - >/dev/null ; done
	@rm -f Gopkg.awk
endif

%.all.json: %.toml
	$(eval lang=$(shell basename `echo $*`))
	$(eval i18n=$(shell dirname `echo $@`))
	@$(GOI18N) -sourceLanguage $(lang) -outdir $(i18n) ./$(i18n)/$(lang).toml
	@$(GOI18N) merge -sourceLanguage $(lang) -outdir $(i18n)  $(foreach p,$(call ReverseWordList,$(SUBPKGS)),./$(p)/$@) $@

help:
	$(info $(S) In order to build and install this package $(PACKAGE), you need to have golang installed.)
	$(info $(S) One can use its own GOPATH environment for buidling this package :)
	$(info $(T)$$ make GOPATH=$${GOPATH})
	$(info $(S) To run the tests, use :)
	$(info $(T)$$ make test)
	$(info $(T) you may use DEPNOLOCK=1 when using VM for any target that needs vendor :)
	$(info $(T)$$ make DEPNOLOCK=1 vendor)
	$(info $(T)One can pass GOTESTOPTS in the form GOTESTOPTS="options" as in:)
	$(info $(T)$$ make GOTESTOPTS="-v" test)
	$(info In order to test and build a sub-package, one needs :)
	$(info $(T)$$ make PACKAGE=$(PKGROOT)/sub-package-name test)
	$(info $(S) One can also launch $$ godoc -http=:unused-port)
	$(info $(T)say unused-port is 6060 then)
	$(info $(T)point your web browser at 'http://localhost:6060/pkg')
	$(info $(T)the list of standard and local packages are displayed.)
	$(info In order to document a sub-package, one needs :)
	$(info $(T)$$ make PACKAGE=$(PKGROOT)/sub-package-name doc)
	$(info $(S) To install the package and sub-packages, use :)
	$(info $(T)$$ make install)
	$(info In order to install a sub-package, one needs :)
	$(info $(T)$$ make PACKAGE=$(PKGROOT)/sub-package-name install)
	$(info $(S) To install only the osgeo.org/proj package, use :)
	$(info $(T)$$ make WITHOUTPKGS=1 install)
	$(info $(T)One can pass GOARCH and GOOS for cross-compiling :)
	$(info $(T)$$ GOARM=7 make GOARCH=arm GOOS=linux install)
	$(info Default target both runs the tests and installs the package)
	@$(MAKE) --silent check-env

# always execute these targets when invoked
.PHONY: all tools test clean build-test doc install help

