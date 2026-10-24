# 23 · 用 GORM 保存和读取设备

阶段：gorm。前置：22。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 这题在完整业务中的位置

存储组织：用 GORM 表达已经掌握的 SQL 业务。 [查看完整业务](../../业务主线.md)。

## 这次做成什么

保持 HTTP 行为，用 GORM 替换设备新增与按 ID 查询的数据库实现。

## 具体需求

沿用 devices 表和四个字段，不嵌入会新增其他字段的 gorm.Model。先操作题目复制出来的 SQLite 文件；使用 GORM 传统 API 的 Create、Where、First、WithContext，检查每次 Error。ErrRecordNotFound 映射为 404；重复 IP 仍为 409。依据《数据库准备》使用纯 Go 驱动。

## 输入与结果

```text
同一套 HTTP 用例，换成 GORM 后结果保持一致
原 SQL 版写入的设备仍可查询
```

## 怎样算完成

- [ ] 18、19、21 的请求合同通过；缺记录不返回空设备加 200。
- [ ] SQL 版创建的表和资料能直接读取；没有悄悄新增替代表。
- [ ] 重复 IP、存储失败与取消请求均正确返回，不吞掉 GORM 的 Error。

环境连接说明见 [数据库准备](../../数据库准备.md)，只阅读本阶段对应段落。

## 本题常用写法（遇到再看）

```go
result := db.WithContext(ctx).Where("id = ?", id).First(&device)
```

- 【库合同】这里 db 是 *gorm.DB；SQL 阶段的 db 是 *sql.DB。变量名相同不表示类型相同。
- 【常见用法】传统 API 的执行错误通过 result.Error 检查；不要因为链式调用没有明显的 err 参数就漏掉错误。
- 【命名惯例】GORM 默认把 Device 映射到 devices，把 ID 识别为主键，把 DeviceName 映射到 device_name；可用显式设置调整。
- 【库合同】First 找不到记录会给 ErrRecordNotFound；检查后才能把 device 当成查到的结果。自动命名不代替你核对真实表结构。

[更多常用写法](../../常用写法速查.md)

## 今天从这里动手

先只替换查询操作；能读取 SQL 版的旧样例后，再替换新增。

只在遇到时查这些名字：gorm.Open、Create、Where、First、WithContext、Error。

做完后，用自己的话回答：**GORM 执行的查询大致对应哪条 SQL？**

[相关文档](https://gorm.io/docs/) · [学习路线](../../README.md) · [上一题](../22-gin-router/README.md) · [下一题](../24-gorm-zero/README.md)
