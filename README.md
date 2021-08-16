### 学习使用go
0. 第一个使用go的程序，不喜勿喷，也欢迎大家给点意见
1. 这是没有系统的学习，慢慢摸索的结果
2. 前端做了手机适应，比较简陋


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


###  登录
![login0](https://user-images.githubusercontent.com/16484266/129593017-3fe21510-036a-49be-8d71-8ca0949f9e10.png)
![login1](https://user-images.githubusercontent.com/16484266/129594376-5a486217-517b-4b00-9cfd-aeaa59f2da14.png)

### 帐号管理
![account0](https://user-images.githubusercontent.com/16484266/129593254-2d8a2cdc-56f4-4bdb-b723-158203c909b7.png)
![account1](https://user-images.githubusercontent.com/16484266/129594309-3a92724f-b944-469b-b406-1f9831db2bc1.png)

### 行为日志
![log0](https://user-images.githubusercontent.com/16484266/129593366-b4a697cd-4ce9-476d-8bc1-c9ce27573bdb.png)
![log1](https://user-images.githubusercontent.com/16484266/129594917-33db28b1-a868-4d2a-bcbd-a68a91b6257b.png)

### home 
![home](https://user-images.githubusercontent.com/16484266/129593461-05c05209-aa22-46a9-8f8f-2495776b7b74.png)

### 折叠菜单
![折叠菜单](https://user-images.githubusercontent.com/16484266/129595062-ace961ca-72a0-487a-b8b8-5b8061af4d78.png)





