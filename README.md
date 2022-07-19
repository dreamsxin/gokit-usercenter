
## 安装 kit-cli
```shell
go install github.com/GrantZheng/kit@latest
```

## 创建服务
```shell
cd dreamsxin
kit new service gokitusercenter --module github.com/dreamsxin/gokitusercenter

# 默认 HTTP
kit g s gokitusercenter
kit g s gokitusercenter -t grpc
```