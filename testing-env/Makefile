MAJOR := 1
MINOR := 0
PATCH := 0
VERSION := ${MAJOR}.${MINOR}.${PATCH}
APPLICATION := testing
NAME = ${APPLICATION}-v${VERSION}

ENV ?= linux
ARCH := amd64
SUPPORTED_OS := linux freebsd windows darwin

BINARY_DIR := bin
BUILD_NAME := $(NAME)-$(ENV)-$(ARCH)
SOURCE_DIR := 

all:ensure default

ensure:
	@mkdir -p $(BINARY_DIR)


default:
	printf("build test bin")
