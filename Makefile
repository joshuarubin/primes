ALL_DIRS=$(shell find . \( -path ./Godeps -o -path ./.git \) -prune -o -type d -print)
GO_PKGS=$(shell go list ./...)
EXECUTABLE=primes
EXECUTABLE_DIR=cmd/$(EXECUTABLE)
GO_FILES=$(foreach dir, $(ALL_DIRS), $(wildcard $(dir)/*.go))
EXECUTABLE_PKG=$(filter %$(EXECUTABLE_DIR), $(GO_PKGS))

ifeq ("$(CIRCLECI)", "true")
	CI_SERVICE = circle-ci
endif

all: build

lint:
	golint ./...
	go vet ./...

test:
	godep go test -v -race ./...

coverage: .acc.out

.acc.out: $(GO_FILES)
	@echo "mode: set" > .acc.out
	@for pkg in $(GO_PKGS); do \
		cmd="godep go test -v -coverprofile=profile.out $$pkg"; \
		eval $$cmd; \
		if test $$? -ne 0; then \
			exit 1; \
		fi; \
		if test -f profile.out; then \
			cat profile.out | grep -v "mode: set" >> .acc.out; \
		fi; \
	done
	@rm -f ./profile.out

coveralls: .coveralls-stamp

.coveralls-stamp: .acc.out
	@if [ -n "$(COVERALLS_REPO_TOKEN)" ]; then \
		goveralls -v -coverprofile=.acc.out -service $(CI_SERVICE) -repotoken $(COVERALLS_REPO_TOKEN); \
	fi
	@touch .coveralls-stamp

build: $(EXECUTABLE)

clean:
	@rm -f ./.acc.out \
      $(EXECUTABLE) \
      ./.coveralls-stamp

save: Godeps/Godeps.json
	
Godeps/Godeps.json: $(GO_FILES)
	@rm -rf ./Godeps
	GOOS=linux GOARCH=amd64 godep save ./...

$(EXECUTABLE): $(GO_FILES)
	godep go build -v -o $(EXECUTABLE) $(EXECUTABLE_PKG)

.PHONY: all lint test coverage coveralls build clean save
