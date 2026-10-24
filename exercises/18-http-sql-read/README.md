# 18 · 通过 HTTP 查询已保存的设备

阶段：HTTP + SQL。前置：13、06；建议完成 17 后进入。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 这题在完整业务中的位置

设备详情接口：把网址中的设备编号连接到真实数据库查询。 [查看完整业务](../../业务主线.md)。

## 这次做成什么

把按 ID 查询接到一个 HTTP 接口。

## 具体需求

GET /devices/{id}：成功 200，返回 {"id":1,"name":"sw-core","ip":"192.0.2.10","enabled":true}。非正整数 ID 为 400 {"error":"invalid request"}；不存在 404 {"error":"device not found"}；数据库故障 500 {"error":"internal error"}。错误详情只记录到本地日志。所有响应为 JSON。

## 输入与结果

```text
GET /devices/1 → 200 + 设备
GET /devices/abc → 400
GET /devices/999999 → 404（前提该 ID 不存在）
```

## 怎样算完成

- [ ] 临时库插入样例，用实际生成的 ID 查询成功。
- [ ] 坏 ID、缺记录、数据库故障分别得到 400/404/500。
- [ ] httptest 测试真实路由，不能只调用处理器而漏掉路径参数。

## 本题常用写法（遇到再看）

当处理器需要持有数据库时，常见的方法声明长这样（先看形状）：

```go
func (h *httpHandler) getDevice(w http.ResponseWriter, r *http.Request) {
    // 通过 h 使用处理器持有的依赖。
}
```

- 【语法】`(h *httpHandler)` 是接收者；h 指向一个处理器对象。它放在方法名前面，与后面的 w、r 参数分开。
- 【习惯】接收者常取类型的短名称 h，并在同类型方法中保持一致；httpHandler 是自定义类型，不是 Go 内置名称。
- 【常见调用】`h.getDevice` 把绑定了 h 的方法交出去；`h.getDevice(w, r)` 才是当场调用。
- 【库合同】处理器仍没有 Go 返回值；通过 w 写响应。数据库函数可以返回 Device、error，由处理器转成 HTTP 结果。普通函数写法也能完成本题。

[更多常用写法](../../常用写法速查.md)

## 今天从这里动手

先把第 13 题的查询提取为独立函数，在处理器里调用；无需先建立多层目录。

只在遇到时查这些名字：r.PathValue、strconv、HTTP 状态码、函数复用。

做完后，用自己的话回答：**从网址里的字符串 ID 到返回 JSON，中间经过哪些步骤？**

[相关文档](https://pkg.go.dev/net/http#Request.PathValue) · [学习路线](../../README.md) · [上一题](../17-sql-transaction/README.md) · [下一题](../19-http-sql-create/README.md)
