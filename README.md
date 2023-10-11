使用 go-zero 官方的工具生成代码：

# 创建 API 服务
    goctl api new core

运行生成的代码：

# 启动服务
    go run core.go -f etc/core-api.yaml

# 使用 api 文件生成代码
    goctl api go -api core.api -dir . -style go_zero

数据来源：https://szgj.2500.tv/
