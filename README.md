# admin
---
### 依赖
- golang 1.6.3
- gin
- jquery
- layui
- font-awesome
- animate

### 初始化数据库
- 查看 init.sql 

### 开放防火墙
```shell
firewall-cmd --add-port=20010/tcp --permanent
firewall-cmd --reload
// 调试使用的端口 
firewall-cmd --add-port=2345/tcp --permanent
```

### 模板重新定义分隔符
```javascript
    laytpl.config({
        open: '<%',
        close: '%>'
    });
```
