.PHONY: help run build test clean docker-build docker-up docker-down install

help: ## 显示帮助信息
	@echo "可用命令:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## 安装依赖
	go mod download
	go mod tidy

run: ## 运行服务
	go run cmd/server/main.go

build: ## 编译项目
	CGO_ENABLED=0 go build -o bin/server cmd/server/main.go

build-linux: ## 编译Linux版本
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server-linux cmd/server/main.go

test: ## 运行所有测试
	go test -v ./...

test-api: ## 运行API测试
	go test -v ./test/api/...

test-coverage: ## 运行测试并生成覆盖率报告
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

clean: ## 清理编译文件
	rm -rf bin/
	rm -rf logs/*.log
	rm -f coverage.out

docker-build: ## 构建Docker镜像
	docker build -t pet-adoption-platform:latest -f docker/Dockerfile .

docker-up: ## 启动Docker容器
	cd docker && docker-compose up -d

docker-down: ## 停止Docker容器
	cd docker && docker-compose down

docker-logs: ## 查看Docker日志
	cd docker && docker-compose logs -f app

fmt: ## 格式化代码
	go fmt ./...
	gofmt -s -w .

lint: ## 代码检查
	golangci-lint run

.DEFAULT_GOAL := help
