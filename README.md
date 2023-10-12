# 后端 API

## 使用方法
1. 进入目录安装依赖
```shell
cd bus_api
go mod tidy
```
2. 复制配置文件
```shell
cp core/etc/core-api.yaml.example core/etc/core-api.yaml
```
3. 导入数据库 database/core.sql

4. 运行服务
```shell
go run core/core.go -f core/etc/core-api.yaml
```
5. 启动后浏览器访问 http://localhost:8899，如果需更改启动端口，可在 core-api.yaml 文件中配置。

## 开发说明
使用 go-zero 官方的工具生成代码：

# 使用 api 文件生成代码
    goctl api go -api core.api -dir . -style go_zero

数据来源：https://szgj.2500.tv/
