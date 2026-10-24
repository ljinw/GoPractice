# 19 · 通过 HTTP 保存设备

阶段：HTTP + SQL。前置：18、12、04。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 这题在完整业务中的位置

设备登记接口：用户提交资料后能再次查到。 [查看完整业务](../../业务主线.md)。

## 这次做成什么

用 POST 保存设备，再用上一题的 GET 读到它。

## 具体需求

POST /devices，输入 {"name":"sw-core","ip":"192.0.2.10"}；成功 201，返回完整 Device（enabled=true）；格式错误、空 name、坏 IP 为 400 {"error":"invalid request"}；重复 IP 为 409 {"error":"ip already exists"}；存储故障为 500 {"error":"internal error"}。使用 HTTP 请求的 Context 执行数据库操作。

## 输入与结果

```text
POST /devices → 201 + 含 id 的设备
GET /devices/返回的id → 相同设备
重复 POST 同 IP → 409
```

## 怎样算完成

- [ ] 新增→查询→重启服务后查询，资料仍在。
- [ ] 无效输入不会新增；重复 IP 不会产生第二条记录。
- [ ] 两个并发请求插入相同 IP：恰好一个 201、一个 409，数据库只多一行；若出现 locked，应先检查 SQLite 连接配置。

并发验收在前两项通过后做。SQLite 学习阶段可设 SetMaxOpenConns(1)，让唯一约束清楚地表现为一个成功、一个冲突；不要用请求前先查一次来替代约束。测试使用同一个服务实例。

## 本题常用写法（遇到再看）

```go
ctx := r.Context()
```

- 【习惯】ctx 放在业务函数的第一个参数，沿调用链传下去；来自请求的取消信号由此传到数据库操作。
- 【常见用法】已经有 r.Context() 时，通常不换成 context.Background()，否则会丢掉这次请求的取消关系。
- 【HTTP 常用表达】`http.StatusCreated` 是 201，`http.StatusConflict` 是 409；本题用它们表达创建成功与 IP 冲突。
- 【常见命名】输入可叫 `CreateDeviceRequest`，输出可叫 `Device`；输入结构体无需让客户端填写服务端生成的 ID。

[更多常用写法](../../常用写法速查.md)

## 今天从这里动手

沿用第 18 题的查询；只补一个 POST 处理器，调用第 12 题的新增函数。

只在遇到时查这些名字：r.Context、201、409、错误分类、数据库唯一约束。

做完后，用自己的话回答：**为什么 HTTP 处理器和数据库新增函数各自负责一部分？**

[相关文档](https://go.dev/doc/database/cancel-operations) · [学习路线](../../README.md) · [上一题](../18-http-sql-read/README.md) · [下一题](../20-chi-router/README.md)
