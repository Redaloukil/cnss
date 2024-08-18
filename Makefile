ENV ?= linux
NAME = main
ARCH := amd64
BINARY_DIR := bin
BUILD_NAME := $(NAME)_$(ENV)_$(ARCH)
SOURCE_DIR := ./cmd/cnss/main.go 

all:ensure 

ensure:
		mkdir -p $(BINARY_DIR)

default:
	@echo -e "BUILD INSTRUCTION\n"

	@echo -e "for building the project"
	@echo -e "run command >> make build ENV={your os linux,freebsd,windows,darwin}"
	@echo -e "for example on mac run >> make build ENV=darwin"
	@echo -e "for testing use build-all\n"


	@echo -e "for running the project"
	@echo -e "run command >> make run ENV={your os linux,freebsd,windows,darwin}"
	@echo -e "for example on mac run >> make build ENV=darwin\n"
  
	@echo -e "for cleaning up use >> make clean"


build:
	@echo -e "\nBuilding for os : $(ENV) $(ARCH) architecture"
	GOOS=$(ENV) GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(BUILD_NAME) $(SOURCE_DIR)
	@echo -e "build completed at $(BINARY_DIR)/$(BUILD_NAME)\n"


build-all:

	@echo -e "\nBuilding for all os : linux,freebsd,windows,darwin $(ARCH) architecture"
	GOOS=linux GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(NAME)-linux-$(ARCH) $(SOURCE_DIR)
	@echo -e "build completed at $(BINARY_DIR)/$(NAME)_linux_$(ARCH)

	GOOS=freebsd GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(NAME)-freebsd-$(ARCH) $(SOURCE_DIR)
	@echo -e "build completed at $(BINARY_DIR)/$(NAME)_freebsd_$(ARCH)

	GOOS=windows GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(NAME)-windows-$(ARCH) $(SOURCE_DIR)
	@echo -e "build completed at $(BINARY_DIR)/$(NAME)_windows_$(ARCH)

	GOOS=darwin GOARCH=$(ARCH) go build -o ./$(BINARY_DIR)/$(NAME)-darwin-$(ARCH) $(SOURCE_DIR)
	@echo -e "build completed at $(BINARY_DIR)/$(NAME)_darwin_$(ARCH)

run:
	./${BINARY_DIR}/${BUILD_NAME}

clean: 
	rm -ir ./${BINARY_DIR}/${NAME}*

$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)

.PHONY: ensure default build build-all run clean
