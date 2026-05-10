# 定义 Go 命令别名
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# 输出可执行文件名称
BINARY_NAME=bin/goframe-vben-run

# 构建命令模板
BUILD_COMMAND=$(GOBUILD) -ldflags "-w -s" -v -o $(BINARY_NAME)

# 默认目标
all: build

# 构建规则
build: build-linux build-win build-mac

# 清理规则
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# 测试规则
test:
	$(GOTEST) -v ./...

# 交叉编译规则
build-mac:
	GOOS=darwin GOARCH=amd64 $(BUILD_COMMAND)
build-linux:
	GOOS=linux GOARCH=amd64 $(BUILD_COMMAND)
build-win:
	GOOS=windows GOARCH=amd64 $(BUILD_COMMAND)
