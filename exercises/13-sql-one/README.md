# 13 · 按编号找一台设备

阶段：database/sql。前置：12。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 这题在完整业务中的位置

定位设备：让后续巡检能够按编号取得设备资料。 [查看完整业务](../../业务主线.md)。

## 这次做成什么

根据 ID 查询一条设备记录，并区分不存在与查询失败。

## 具体需求

ID 为正整数。用 QueryRowContext 和 Scan 读取 id/name/ip/enabled；不存在返回可识别的 not found；数据库故障返回真实错误，不能伪装成不存在。一次查询保持只读。

## 输入与结果

```text
查询已插入 ID → 对应 Device
查询不存在 ID → not found
```

## 怎样算完成

- [ ] 用上一题返回的 ID 查到原数据。
- [ ] 不存在 ID 与关闭数据库句柄后的查询错误能够区分。
- [ ] 0 或负 ID 拒绝；查询前后记录数量一致。

## 本题常用写法（遇到再看）

```go
err := db.QueryRowContext(ctx, query, id).Scan(&device.Name)
```

这是只查询一列 name 的局部示意；查询多列时，Scan 的目标数量和顺序也要匹配。

- 【库合同】QueryRowContext 返回 Row，查询与扫描中的错误通常在 Scan 时拿到；不能仿照 QueryContext 接两个返回值。
- 【常见用法】用 `errors.Is(err, sql.ErrNoRows)` 判断“没有记录”，再处理其他错误。
- 【习惯】`device` 用单数，表示一台；`devices` 常表示列表。
- 【语法与用途】`&device.Name` 让 Scan 能填入字段；Scan 不会自动根据变量名挑选数据库列。

[更多常用写法](../../常用写法速查.md)

## 今天从这里动手

只查 name 一列并扫描到字符串，成功后再扩展到四列。

只在遇到时查这些名字：SELECT、WHERE、QueryRowContext、Scan、sql.ErrNoRows。

做完后，用自己的话回答：**查不到和查坏了，调用者应如何区分？**

[相关文档](https://go.dev/doc/database/querying) · [学习路线](../../README.md) · [上一题](../12-sql-insert/README.md) · [下一题](../14-sql-list/README.md)
