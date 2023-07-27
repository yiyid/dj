BINARY_NAME := deploy-jenkins

# 设置默认目标
all: build

# 定义源文件目录（当前目录）
SRC_FILES := .

# 定义交叉编译的目标操作系统
PLATFORMS := linux windows darwin

# 使用当前操作系统作为默认的构建平台
OS ?= $(shell go env GOOS)

# 使用当前架构作为默认的构建架构
ARCH ?= $(shell go env GOARCH)

# 设置构建输出目录基于操作系统
OUTPUT_DIR := build

# 当前平台的构建目标
build:
	go build -ldflags "-w -s" -o $(BINARY_NAME) $(SRC_FILES)

# 交叉编译的构建目标
# 交叉编译的构建目标
# 交叉编译的构建目标
build-all:
	@for os in $(PLATFORMS); do \
		if [ "$$os" = "windows" ]; then \
			for arch in arm64 amd64 386; do \
				echo "Building $(BINARY_NAME) for $$os/$$arch"; \
				GOOS=$$os GOARCH=$$arch go build -ldflags "-w -s" -o $(OUTPUT_DIR)/$(BINARY_NAME)_$${os}_$$arch $(SRC_FILES); \
			done; \
		else \
			for arch in arm64 amd64; do \
				echo "Building $(BINARY_NAME) for $$os/$$arch"; \
				GOOS=$$os GOARCH=$$arch go build -ldflags "-w -s" -o $(OUTPUT_DIR)/$(BINARY_NAME)_$${os}_$$arch $(SRC_FILES); \
			done; \
		fi \
	done




# 清理目标
clean:
	go clean
	rm -rf $(BINARY_NAME) build/

# 运行目标
run:
	go run $(SRC_FILES)

# 帮助目标
help:
	@echo "使用说明:"
	@echo "make             编译项目（针对当前平台）"
	@echo "make build-all   交叉编译项目到不同操作系统和架构"
	@echo "make clean       清理项目"
	@echo "make run         运行项目"
	@echo "make help        显示帮助信息"
