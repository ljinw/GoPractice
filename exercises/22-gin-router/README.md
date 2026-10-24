# 22 · 用 gin 写同一组接口

阶段：gin。前置：21。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 这题在完整业务中的位置

框架比较：用相同设备业务理解 gin 的输入输出处理。 [查看完整业务](../../业务主线.md)。

## 这次做成什么

在独立副本使用 gin 接入相同的设备新增与查询能力。

## 具体需求

复用数据库与校验函数；将输入/输出适配为 gin.Context 的 Param、ShouldBindJSON、JSON。延续 18、19 的状态码和 JSON 字段；将第 21 题编号中间件适配到 gin。启用方法不匹配的 405 行为，明确设置监听 127.0.0.1。

## 输入与结果

```text
POST /devices → 与 chi 版相同的 201/400/409
GET /devices/{id} → 与 chi 版相同的 200/400/404
```

## 怎样算完成

- [ ] 沿用 18~21 的 HTTP 行为测试；比较 JSON 数据，不比较字段顺序。
- [ ] 绑定失败后只输出一次 400；空 name 仍按业务规则拒绝。
- [ ] 未知路径 404、错误方法 405，也有 X-Request-ID；多次请求之间无状态串用。

## 本题常用写法（遇到再看）

```go
func getDevice(c *gin.Context) {
    // 使用 c 读取输入、组织输出。
}
```

- 【习惯】c 常指 gin.Context；ctx 常指标准库 context.Context，二者不要凭名字混用。
- 【库合同】gin 处理器接收一个 *gin.Context；底层 HTTP 请求在 c.Request，需要标准上下文时取 c.Request.Context()。
- 【常见分工】c.Param 读路径参数，c.Query 读查询参数，c.JSON 写 JSON；名称和参数顺序由库定义。
- 【控制流程】Abort 阻止后续处理器，但不会自动结束当前函数；当前错误分支通常仍需 return。普通处理器中的 return 也不等于中间件链的 Abort。

[更多常用写法](../../常用写法速查.md)

## 今天从这里动手

先只改 GET 查询接口，测试通过再迁移 POST；两步分别验收。

只在遇到时查这些名字：gin.New、gin.Context、ShouldBindJSON、c.JSON、c.Next。

做完后，用自己的话回答：**gin.Context 中哪些操作对应你已学过的 r 和 w？**

[相关文档](https://gin-gonic.com/en/docs/) · [学习路线](../../README.md) · [上一题](../21-chi-request-id/README.md) · [下一题](../23-gorm-store/README.md)
