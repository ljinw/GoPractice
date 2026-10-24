# 03 · 把设备信息返回为 JSON

阶段：net/http。前置：02。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

本阶段入口：[标准库路线](../../第一阶段标准库学习路线.md)。先读本题方法/类型详解，再做需求；新API不靠猜字段或只抄调用。

## 这题在完整业务中的位置

设备详情：把设备资料按约定格式交给调用方。 [查看完整业务](../../业务主线.md)。

## 这次做成什么

把一台写在程序里的假设备变成 JSON 响应。

## 具体需求

新增 GET /device，固定返回 {"name":"sw-core","ip":"192.0.2.10"}。使用结构体和 encoding/json；响应 Content-Type 为 application/json，可带 charset。IP 只是样例数据，不发起网络连接。

## 输入与结果

```text
GET /device → 200
{"name":"sw-core","ip":"192.0.2.10"}
```

## 怎样算完成

- [ ] JSON 能解析，name 与 ip 的值正确；字段顺序不作为判断条件。
- [ ] Content-Type 包含 application/json。
- [ ] /health、/hello 的行为保持一致。

## 本题常用写法（遇到再看）

```go
type Device struct {
    Name string `json:"name"`
}
```

- 【语法】`Name` 首字母大写是导出字段，标准 JSON 编码器能访问；小写 name 不是同一个效果。
- 【库合同】标签 `json:"name"` 指定 JSON 中的名字；反引号是 Go 字符串语法，里面的 json 键由库识别。
- 【常见用法】`w.Header().Set("Content-Type", "application/json")` 要放在提交响应前；它设置头部，不会直接返回 JSON 正文。
- 【习惯】使用 `http.StatusOK` 表达 200，比到处写裸数字更容易读。设备结构体叫 Device 是业务命名选择。

[更多常用写法](../../常用写法速查.md)

## 今天从这里动手

用中文写下设备有哪两个字段，再尝试声明 Device 结构体。

只在遇到时查这些名字：struct、json 标签、json.NewEncoder、w.Header().Set。

做完后，用自己的话回答：**结构体字段名和最终 JSON 字段名为什么可能不同？**

[相关文档](https://pkg.go.dev/encoding/json) · [学习路线](../../README.md) · [上一题](../02-http-query/README.md) · [下一题](../04-http-post/README.md)
