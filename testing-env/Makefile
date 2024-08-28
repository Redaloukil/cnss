MAJOR := 1
MINOR := 0
PATCH := 0
VERSION := v${MAJOR}.${MINOR}.${PATCH}

ENV ?= linux
ARCH := amd64
SUPPORTED_OS := linux freebsd windows darwin
BINARY_DIR := ./bin

APPLICATION_NAME := application
APPLICATION_SOURCE := ./application/main.go
APPLICATION_BUILD := ${BINARY_DIR}/application/${APPLICATION_NAME}-${VERSION}-${ENV}-${ARCH}


SERVER_NAME := server
SERVER_SOURCE := ./server/main.go
SERVER_BUILD := ${BINARY_DIR}/server/${SERVER_NAME}-${VERSION}-${ENV}-${ARCH}


SIMULATOR_NAME := simulator
SIMULATOR_SOURCE := ./simulator/main.go
SIMULATOR_BUILD := ${BINARY_DIR}/simulator/${SIMULATOR_NAME}-${VERSION}-${ENV}-${ARCH}



all:ensure default

ensure:
	@mkdir -p ${BINARY_DIR}
	@mkdir -p ${BINARY_DIR}/${APPLICATION_NAME}
	@mkdir -p ${BINARY_DIR}/${SERVER_NAME}
	@mkdir -p ${BINARY_DIR}/${SIMULATOR_SOURCE}


default:
	@printf "\033[1;34mBUILD INSTRUCTION\033[0m\n\n"

	@printf "🌸 \033[1;32mTo build the testing environments:\033[0m\n"
	@printf "   \033[1;35m>> make build ENV={your os linux, freebsd, windows, darwin}\033[0m\n"
	@printf "   \033[1;36mExample:\033[0m on macOS\n   \033[1;35m>> make build ENV=darwin\033[0m\n\n"

	@printf "🌸 \033[1;32mTo run the testing environments:\033[0m\n"
	@printf "   \033[1;35m>> make run ENV={your os linux, freebsd, windows, darwin}\033[0m\n"
	@printf "   \033[1;36mExample:\033[0m on macOS\n   \033[1;35m>> make run ENV=darwin\033[0m\n\n"

	@printf "🌸 \033[1;31mTo clean up the testing environments:\033[0m\n"
	@printf "   \033[1;35m>> make clean\033[0m\n\n"

build:
	@printf "\033[1;32mBuilding ${SIMULATOR_NAME} for os: $(ENV) $(ARCH) architecture\033[0m\n"
	GOOS=$(ENV) GOARCH=$(ARCH) go build -o ${SIMULATOR_BUILD} ${SIMULATOR_SOURCE}
	@printf "\033[1;32mBuild completed at ${SIMULATOR_BUILD}\033[0m\n\n"
	

build-server:

	@printf "\033[1;32mBuilding ${SERVER_NAME} for os: linux $(ARCH) architecture\033[0m\n"
	GOOS=linux GOARCH=$(ARCH) go build -o ${BINARY_DIR}/server/${server_NAME}-${VERSION}-linux-${ARCH} ${SERVER_SOURCE}
	@printf "\033[1;32mBuild completed at ${server_BUILD}\033[0m\n\n"
	

	@printf "\033[1;32mBuilding ${SERVER_NAME} for os: freebsd $(ARCH) architecture\033[0m\n"
	GOOS=freebsd GOARCH=$(ARCH) go build -o ${BINARY_DIR}/server/${server_NAME}-${VERSION}-freebsd-${ARCH} ${SERVER_SOURCE}
	@printf "\033[1;32mBuild completed at ${SIMULATOR_BUILD}\033[0m\n\n"


	@printf "\033[1;32mBuilding ${SERVER_NAME} for os: darwin $(ARCH) architecture\033[0m\n"
	GOOS=darwin GOARCH=$(ARCH) go build -o ${BINARY_DIR}/server/${server_NAME}-${VERSION}-darwin-${ARCH} ${SERVER_SOURCE}
	@printf "\033[1;32mBuild completed at ${SERVER_BUILD}\033[0m\n\n"

	@printf "\033[1;32mBuilding ${SERVER_NAME} for os: windows $(ARCH) architecture\033[0m\n"
	GOOS=windows GOARCH=$(ARCH) go build -o ${BINARY_DIR}/server/${server_NAME}-${VERSION}-windows-${ARCH}.exe ${SERVER_SOURCE}
	@printf "\033[1;32mBuild completed at ${SERVER_BUILD}\033[0m\n\n"


build-application:

	@printf "\033[1;32mBuilding ${APPLICATION_NAME} for os: linux $(ARCH) architecture\033[0m\n"
	GOOS=linux GOARCH=$(ARCH) go build -o ${BINARY_DIR}/server/${APPLICATION_NAME}-${VERSION}-linux-${ARCH} ${SERVER_SOURCE}
	@printf "\033[1;32mBuild completed at ${APPLICATION_BUILD}\033[0m\n\n"
	

	@printf "\033[1;32mBuilding ${APPLICATION_NAME} for os: freebsd $(ARCH) architecture\033[0m\n"
	GOOS=freebsd GOARCH=$(ARCH) go build -o ${BINARY_DIR}/server/${APPLICATION_NAME}-${VERSION}-freebsd-${ARCH} ${SERVER_SOURCE}
	@printf "\033[1;32mBuild completed at ${APPLICATION_BUILD}\033[0m\n\n"


	@printf "\033[1;32mBuilding ${APPLICATION_NAME} for os: darwin $(ARCH) architecture\033[0m\n"
	GOOS=darwin GOARCH=$(ARCH) go build -o ${BINARY_DIR}/server/${APPLICATION_NAME}-${VERSION}-darwin-${ARCH} ${SERVER_SOURCE}
	@printf "\033[1;32mBuild completed at ${APPLICATION_BUILD}\033[0m\n\n"

	@printf "\033[1;32mBuilding ${APPLICATION_NAME} for os: windows $(ARCH) architecture\033[0m\n"
	GOOS=windows GOARCH=$(ARCH) go build -o ${BINARY_DIR}/server/${APPLICATION_NAME}-${VERSION}-windows-${ARCH}.exe ${SERVER_SOURCE}
	@printf "\033[1;32mBuild completed at ${APPLICATION_BUILD}\033[0m\n\n"



run: build
	@printf "\n🌸 \033[1;32mExecuting ${APPLICATION_BUILD}....\033[0m\n"
	@./${APPLICATION_BUILD}

	@printf "\n🌸 \033[1;32mExecuting ${SERVER_BUILD}....\033[0m\n"
	@./${SERVER_BUILD}

	@printf "\n🌸 \033[1;32mExecuting ${SIMULATOR_BUILD}....\033[0m\n"
	@./${SIMULATOR_BUILD}


clean: 
	@printf "\033[1;31m" 
	rm -ir ${BINARY_DIR}/*/*
	@printf "\033[1;32mbinary cleaned from ${BINARY_DIR}/ \033[0m\n\n" 
	

.PHONY: ensure all default build run clean
