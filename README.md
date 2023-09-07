# douyin

原仓库地址(也是我组同学的仓库): 
https://github.com/Diuuhiy/douyin

运行说明:

1. 启动 mysql etcd server

2. 启动 rpc 服务
```bash
cd rpc/core
go run user.go -f etc/user.yaml
cd rpc/interactive
go run interactive.go -f etc/interactive.yaml
cd rpc/social
go run social.go -f etc/social.yaml
```

3. 启动 api
```bash
cd api/douyin
go run douyin.go -f etc/douyin.yaml
```

