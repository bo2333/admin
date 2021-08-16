---
### 依赖
- golang 1.6.3
- gin
- jquery
- layui
- font-awesome
- MariaDB

### 初始化数据库
- 查看 init.sql 

### 开放防火墙
```shell
firewall-cmd --add-port=1999/tcp --permanent
firewall-cmd --reload
// 调试使用端口 
firewall-cmd --add-port=2345/tcp --permanent
```

### 重新定义的分隔符
```
router.Delims("{[", "]}")
```


### 步骤
1. go mod init 
2. 安装MariaDB数据库, 建立admin库
3. 数据库运行init.sql 
4. 进行配置 resources/conf/config.ini
5. 运行 go run main.go 启动服务
6. 127.0.0.1:1999/login 登录