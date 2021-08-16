### 学习使用go
0. 第一个使用go的程序，不喜勿喷，也欢迎大家给点意见
1. 这是没有系统的学习，慢慢摸索的结果


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

### 图片展
0. 简单的做了手机适应，做工一般，希望大家能给点意见。
1. 登录
![image](https://user-images.githubusercontent.com/16484266/129593017-3fe21510-036a-49be-8d71-8ca0949f9e10.png)

2. 帐号管理
![1629129510(1)](https://user-images.githubusercontent.com/16484266/129593254-2d8a2cdc-56f4-4bdb-b723-158203c909b7.png)

3. 行为日志
![1629129566(1)](https://user-images.githubusercontent.com/16484266/129593366-b4a697cd-4ce9-476d-8bc1-c9ce27573bdb.png)

4. home 
![1629129613(1)](https://user-images.githubusercontent.com/16484266/129593461-05c05209-aa22-46a9-8f8f-2495776b7b74.png)

