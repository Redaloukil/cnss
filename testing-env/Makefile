# NOTE :-
# MAJOR => 1 for breking changes
# MINOR => 0 for new feature
# PATCH => 0 for bug patch/minor fix

ARCH := amd64

APPLICATION_VERSION := 1.0.0
SERVER_VERSION := 1.0.0
SIMULATOR_VERSION := 1.0.0

# colour 
Red    := \033[1;31m
Green  := \033[1;32m
Yellow := \033[1;33m
Blue   := \033[1;34m
Purple := \033[1;35m
Cyan   := \033[1;36m
Clear  := \033[0m

OS :=

ifeq ($(OS),Windows_NT)
	OS := windows
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OS := linux
	endif
	ifeq ($(UNAME_S),Darwin)
		OS := darwin
	endif
	ifeq ($(UNAME_S),FreeBSD)
		OS := FreeBSD
	endif
	# fllback
	ifeq ($(OS),)
		OS := linux
	endif
endif


default: help

ensure:
	@mkdir -p ./bin
	@mkdir -p ./bin/server/
	@mkdir -p ./bin/application/
	@mkdir -p ./bin/simulator/

help:

	@printf "$(Cyan)Usage: make [target]$(Clear)\n\n"
	@printf "$(Blue)Targets:$(Clear)\n"
	@printf "$(Purple)   help$(Clear)            - Display this help message\n"
	@printf "$(Purple)   ensure$(Clear)          - Create necessary directories\n"
	@printf "$(Purple)   build$(Clear)           - Build the Server, Application, and Simulator for the current OS and architecture\n"
	@printf "$(Purple)   build-all$(Clear)       - Build the Server, Application, and Simulator for all supported OSes (linux, darwin, freebsd, windows)\n"
	@printf "$(Purple)   detect-os$(Clear)       - Default OS detected for build target\n"
	@printf "$(Purple)   run-server$(Clear)      - Run the Server for the current OS and architecture\n"
	@printf "$(Purple)   run-application$(Clear) - Run the Application for the current OS and architecture\n"
	@printf "$(Purple)   run-simulator$(Clear)   - Run the Simulator for the current OS and architecture\n"
	@printf "$(Purple)   clean$(Clear)           - Remove all built binaries\n"

	@printf "\n$(Red)Example:$(Clear)\n"
	@printf "$(Yellow)>> make detected-os $(Clear)\n"
	@$(MAKE) --no-print-directory detect-os 


detect-os:
	@printf "\n$(Green)Your Detected OS is $(OS)$(Clear)\n"


build: ensure

	@printf "$(Purple) building Server for default OS $(OS) $(Clear)\n"
	go build -o ./bin/server/server-v${SERVER_VERSION}-${OS}-${ARCH} ./server/main.go
	@printf "$(Purple) building Application for default OS $(OS) $(Clear)\n"
	go build -o ./bin/application/application-v${APPLICATION_VERSION}-${OS}-${ARCH} ./application/main.go
	@printf "$(Purple) building Simmulator for default OS $(OS) $(Clear)\n"
	go build -o ./bin/simulator/simulator-v${SIMULATOR_VERSION}-${OS}-${ARCH} ./simulator/main.go

build-all: ensure

	@printf "$(Purple) building Server, Application, & Simmulator for linux $(Clear)\n"
	GOOS=linux GOARCH=$(ARCH) go build -o ./bin/server/server-v${SERVER_VERSION}-linux-${ARCH} ./server/main.go
	GOOS=linux GOARCH=$(ARCH) go build -o ./bin/application/application-v${APPLICATION_VERSION}-linux-${ARCH} ./application/main.go
	GOOS=linux GOARCH=$(ARCH) go build -o ./bin/simulator/simulator-v${SIMULATOR_VERSION}-linux-${ARCH} ./simulator/main.go

	@printf "$(Purple) building Server, Application, & Simmulator for freebsd $(Clear)\n"
	GOOS=freebsd GOARCH=$(ARCH) go build -o ./bin/server/server-v${SERVER_VERSION}-freebsd-${ARCH} ./server/main.go
	GOOS=freebsd GOARCH=$(ARCH) go build -o ./bin/application/application-v${APPLICATION_VERSION}-freebsd-${ARCH} ./application/main.go
	GOOS=freebsd GOARCH=$(ARCH) go build -o ./bin/simulator/simulator-v${SIMULATOR_VERSION}-freebsd-${ARCH} ./simulator/main.go

	@printf "$(Purple) building Server, Application, & Simmulator for darwin $(Clear)\n"
	GOOS=darwin GOARCH=$(ARCH) go build -o ./bin/server/server-v${SERVER_VERSION}-darwin-${ARCH} ./server/main.go
	GOOS=darwin GOARCH=$(ARCH) go build -o ./bin/application/application-v${APPLICATION_VERSION}-darwin-${ARCH} ./application/main.go
	GOOS=darwin GOARCH=$(ARCH) go build -o ./bin/simulator/simulator-v${SIMULATOR_VERSION}-darwin-${ARCH} ./simulator/main.go

	@printf "$(Purple) building Server, Application, & Simmulator for windows $(Clear)\n"
	GOOS=windows GOARCH=$(ARCH) go build -o ./bin/server/server-v${SERVER_VERSION}-windows-${ARCH} ./server/main.go
	GOOS=windows GOARCH=$(ARCH) go build -o ./bin/application/application-v${APPLICATION_VERSION}-windows-${ARCH} ./application/main.go
	GOOS=windows GOARCH=$(ARCH) go build -o ./bin/simulator/simulator-v${SIMULATOR_VERSION}-windows-${ARCH} ./simulator/main.go
	
run-server: 
	@printf "$(Purple) running Server $(Clear)\n"
	./bin/server/server-v1.0.0-${OS}-amd64

run-application: 
	@printf "$(Purple) running Application $(Clear)\n"
	./bin/application/application-v1.0.0-${OS}-amd64

run-simulator: 
	@printf "$(Purple) running Simmulator $(Clear)\n"
	./bin/simulator/simulator-v1.0.0-${OS}-amd64

# run-all: 
# 	@$(MAKE) run-server &
# 	@$(MAKE) run-application &
# 	@$(MAKE) run-simulator &
# 	wait
# 	@printf "$(Purple) All services started $(Clear)\n"

clean:
	@printf "$(Red)" 
	rm -ir ./bin/*/*
	@printf "$(Green) binary cleaned from ./bin/ $(Clear)\n" 

.PHONY: default ensure detect-os help build build-all run-server run-application run-simulator clean
