MAJOR := 1
MINOR := 0
PATCH := 0
VERSION := ${MAJOR}.${MINOR}.${PATCH}
NAME = main${VERSION}

ENV ?= linux
ARCH := amd64
SUPPORTED_OS := linux freebsd windows darwin

BINARY_DIR := bin
BUILD_NAME := $(NAME)_$(ENV)_$(ARCH)
SOURCE_DIR := ./cmd/cnss/main.go 


all:ensure default

ensure:
	@mkdir -p $(BINARY_DIR)


default:
	@printf "\033[1;34mBUILD INSTRUCTION\033[0m\n\n"

	@printf "🌸 \033[1;32mTo build the project:\033[0m\n"
	@printf "   \033[1;35m>> make build ENV={your os linux, freebsd, windows, darwin}\033[0m\n"
	@printf "   \033[1;36mExample:\033[0m on macOS\n   \033[1;35m>> make build ENV=darwin\033[0m\n\n"
	@printf "🌸 \033[1;32mTo build for all environments:\033[0m\n"
	@printf "   \033[1;35m>> make build-all\033[0m\n\n"

	@printf "🌸 \033[1;32mTo run the project:\033[0m\n"
	@printf "   \033[1;35m>> make run ENV={your os linux, freebsd, windows, darwin}\033[0m\n"
	@printf "   \033[1;36mExample:\033[0m on macOS\n   \033[1;35m>> make run ENV=darwin\033[0m\n\n"

	@printf "🌸 \033[1;31mTo clean up:\033[0m\n"
	@printf "   \033[1;35m>> make clean\033[0m\n\n"


build:
	@printf "\033[1;32mBuilding for os: $(ENV) $(ARCH) architecture\033[0m\n"
	GOOS=$(ENV) GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(BUILD_NAME) $(SOURCE_DIR)
	@printf "\033[1;32mBuild completed at $(BINARY_DIR)/$(BUILD_NAME)\033[0m\n\n"


build-all:
	@printf "\033[1;32mBuilding for all OS: $(SUPPORTED_OS) $(ARCH) architecture\033[0m\n\n"
	GOOS=linux GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(NAME)_linux_$(ARCH) $(SOURCE_DIR)
	@printf "\033[1;32mBuild completed at $(BINARY_DIR)/$(NAME)_linux_$(ARCH)\033[0m\n\n"

	GOOS=freebsd GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(NAME)_freebsd_$(ARCH) $(SOURCE_DIR)
	@printf "\033[1;32mBuild completed at $(BINARY_DIR)/$(NAME)_freebsd_$(ARCH)\033[0m\n\n"

	GOOS=windows GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(NAME)_windows_$(ARCH).exe $(SOURCE_DIR)
	@printf "\033[1;32mBuild completed at $(BINARY_DIR)/$(NAME)_windows_$(ARCH)\033[0m\n\n"

	GOOS=darwin GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(NAME)_darwin_$(ARCH) $(SOURCE_DIR)
	@printf "\033[1;32mBuild completed at $(BINARY_DIR)/$(NAME)_darwin_$(ARCH)\033[0m\n\n" 


run: build
	@printf "🌸 \033[1;32mExecuting ./${BINARY_DIR}/${BUILD_NAME}....\033[0m\n"
	@./${BINARY_DIR}/${BUILD_NAME}


clean: 
	@printf "\033[1;31m" 
	rm -ir ./${BINARY_DIR}/${NAME}*
	@printf "\033[1;32mbinary cleaned from ./${BINARY_DIR}/ \033[0m\n\n" 


.PHONY: ensure all default build build-all run clean
