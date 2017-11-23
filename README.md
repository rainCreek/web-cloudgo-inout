# web-cloudgo-inout
一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）


## 基本要求
1. 支持静态文件服务
2. 支持简单 js 访问
3. 提交表单，并输出一个表格
4. 对`` /unknown ``给出开发中的提示，返回码`` 5xx``

## 运行方法
* 用浏览器访问 ``http://localhost:8080``，显示网页静态网页，支持简答js访问
* 用浏览器访问 ``http://localhost:8080/fillin``，可填写个人信息，并打印输出相应表格
* 用浏览器访问 ``http://localhost:8080/unknown``，返回 ``5xx : Not Implement!Unknown Page! ``


## 实现要点
* 动态文件使用http.HandleFunc进行设置，静态文件使用http.FileServer
* 使用template.ParseFile模版文件解析
* handler处理对应相应的一级路径

## 目录

![](http://img.blog.csdn.net/20171121214306511)
    
## 运行截图
**1. 用浏览器访问 ``http://localhost:8080``，显示网页静态网页**
  
>  启动main.go

![](http://img.blog.csdn.net/20171121213304774)
  
  
> 用浏览器访问 ``http://localhost:8080``
出现以下简单网页。
状态码200：正常。

![](http://img.blog.csdn.net/20171121213329641)
  
  
> curl测试

![](http://img.blog.csdn.net/20171123205714828)

**2. 用浏览器访问 ``http://localhost:8080/fillin``

> 用浏览器访问 ``http://localhost:8080/fillin``
出现信息采集页面。可输入用户信息。
状态码200：正常。

![](http://img.blog.csdn.net/20171121213547043)
  
> 输入用户信息，点击“ok”，进入表格输出页面，打印刚刚输入的用户信息。

![](http://img.blog.csdn.net/20171121213718977)
![](http://img.blog.csdn.net/20171121213731816)

> curl测试
![](http://img.blog.csdn.net/20171121213851854)


**3. 用浏览器访问 ``http://localhost:8080/unknown``

> 用浏览器访问 ``http://localhost:8080/unknown``
网页提示：``5xx : Not Implement!Unknown Page! ``。

![](http://img.blog.csdn.net/20171123210252927)
  
  
> curl测试

![](http://img.blog.csdn.net/20171123205725227)



  
  
