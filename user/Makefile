
# 项目变量定义
API_FILE := user.api
MAIN_FILE := ./user.go
SERVICE_NAME := user

# 构建输出目录
OUTPUT_DIR := ./bin
DOC_DIR := ./doc


.PHONY: format
format:
	@goctl api format --dir .
	@go fmt ./...

# 生成API文档
.PHONY: doc
doc:
	@mkdir -p $(DOC_DIR)
	@rm -rf $(DOC_DIR)/$(SERVICE_NAME).md
	@rm -rf $(DOC_DIR)/$(SERVICE_NAME).json
	@goctl api doc --dir . --o $(DOC_DIR)
	@goctl api plugin -plugin goctl-swagger="swagger -filename $(SERVICE_NAME).json" -api $(API_FILE) -dir $(DOC_DIR)


# 生成Go-Zero代码
.PHONY: goctl.gen
goctl.gen: doc
	goctl api go --dir . --api $(API_FILE)
	go mod tidy

# 初始化项目依赖
.PHONY: init
init:
	go install github.com/zeromicro/go-zero/tools/goctl@latest
	go install github.com/zeromicro/goctl-swagger@latest
	go install github.com/air-verse/air@latest
	goctl env check --install --verbose --force
	go mod tidy

# 构建服务
.PHONY: build
build:
	go build .

# 开发
.PHONY: dev
dev:
	air -c .air.toml

# 一键式开发环境设置
.PHONY: dev-setup
dev-setup: init goctl.gen
