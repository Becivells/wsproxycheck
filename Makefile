.DEFAULT_GOAL:=help

COLOR_ENABLE=$(shell tput colors > /dev/null; echo $$?)
ifeq "$(COLOR_ENABLE)" "0"
CRED:=$(shell tput setaf 1 2>/dev/null)
CGREEN:=$(shell tput setaf 2 2>/dev/null)
CYELLOW:=$(shell tput setaf 3 2>/dev/null)
CBLUE:=$(shell tput setaf 4 2>/dev/null)
CMAGENTA:=$(shell tput setaf 5 2>/dev/null)
CCYAN:=$(shell tput setaf 6 2>/dev/null)
CWHITE:=$(shell tput setaf 7 2>/dev/null)
CEND:=$(shell tput sgr0 2>/dev/null)
endif
.PHONY: build
build:      	## build snapshot
	@echo "$(CGREEN)build snapshot no publish ...$(CEND)"
	@goreleaser build --snapshot --rm-dist

.PHONY: snapshot
snapshot:    	## pre snapshot
	@echo "$(CGREEN)release snapshot no publish ...$(CEND)"
	@goreleaser release --skip-publish  --snapshot --rm-dist
.PHONY: release
release:		## release no publish
	@echo "$(CGREEN)release no publish ...$(CEND)"
	@goreleaser release --skip-publish  --rm-dist

.PHONY: clean
clean:      	## clean up
	@echo "$(CGREEN)clean up dist ...$(CEND)"
	@rm -rf ./dist
.PHONY: help
help:			## Show this help.
	@echo "$(CGREEN)project$(CEND)"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make $(CYELLOW)<target>$(CEND) (default: help)\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  $(CCYAN)%-12s$(CEND) %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
