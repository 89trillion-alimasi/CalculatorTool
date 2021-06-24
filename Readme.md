# Go技术文档大纲-计算器

## 1.整体框架

​	使用Gin框架完成一个接口：实现了简单计算器，通过http发送表达式请求，来进行后端计算在返回值。流程图叫processon 和Readme.md同路径

## 2.目录结构

```
.
├── Readme.md
├── app
│   ├── main
│   └── main.go #程序运行启动
├── ctr
│   └── controller.go 
├── go.mod
├── go.sum
├── model
│   └── model.go #数据结构
├── processon.png #流程图
├── router
│   └── router.go #路由
├── service
│   ├── service.go #计算器业务逻辑
│   └── service_test.go #单元测试
└── test
    ├── __pycache__
    │   └── locust_test.cpython-39.pyc
    ├── cal_test_report.html #压力测试报告
    └── locust_test.py #压力测试


```



## 3.代码逻辑分层

​	

| 层        | 文件夹              | 主要职责               | 调用关系                 | 其他说明 |
| --------- | ------------------- | ---------------------- | ------------------------ | -------- |
| 应用层    | /app/main.go        | 程序启动               | 调用路由层和工具层       |          |
| 路由层    | /router/router.go   | 路由的初始化，路由转发 | 调用控制层               |          |
| 控制层    | /ctr/controller     | 处理请求和构建回复消息 | 被路由层调用，调用服务层 |          |
| service层 | /service/service.go | 计算器业务逻辑         | 被控制层调用             |          |
| model层   | /model/model.go     | 设计数据结构           | 被服务层所调用           |          |



## 4.存储设计

​	

| 内容 | field | 类型  |
| ---- | ----- | ----- |
| 栈   | Stack | []int |



## 5.接口设计

### 	请求方式

​		http get

### 	接口地址

​		http://127.0.0.1:8080/calculator?expr=3+5/2

### 	请求参数

​		例如expr=3+5/2

### 	响应状态吗	

| 状态码 | 说明         |
| ------ | ------------ |
| 200    | 成功         |
| 400    | 返回错误信息 |

### 	

## 6.第三方库

​	无

## 7.如何编译执行

​	进入app 执行./main 可执行文件再打开http://127.0.0.1:8080/calculator?expr=“  ”在引号位置中加入表达式

## 8.todo

1.后期尝试用语法分析器来检查表达式合法性

2.使用逆波兰来计算表达式

3.后续拓展% ^等运算符