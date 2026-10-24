# 14 · 列出全部设备

阶段：database/sql。前置：13。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 这题在完整业务中的位置

设备清单：给批量操作准备可查询的设备集合。 [查看完整业务](../../业务主线.md)。

## 这次做成什么

读取多行设备，按 ID 升序返回。

## 具体需求

用 QueryContext 查询全部设备，ORDER BY id ASC。遍历 rows，逐行 Scan，检查 rows.Err 并关闭 rows。空表返回长度为 0 的列表，不返回 not found。暂不做分页与模糊搜索。

## 输入与结果

```text
空表 → []
三台设备 → 按 id 升序的三个 Device
```

## 怎样算完成

- [ ] 空表正常返回空列表。
- [ ] 插入三条样例后顺序稳定、字段完整。
- [ ] 查询故障向上返回；不能将半份结果当完整成功。

## 本题常用写法（遇到再看）

```go
rows, err := db.QueryContext(ctx, query)
```

- 【习惯】`rows` 指结果集；`row` 常指单行。名字不能替代对真实返回类型的检查。
- 【常见顺序】先检查 err，再安排 rows.Close；Next 为 true 时 Scan，循环后检查 rows.Err。
- 【常见用法】`devices := make([]Device, 0)` 表达一个已初始化的空列表；标准 JSON 编码时它是 []，nil 切片通常编码为 null。
- 【SQL 合同】需要稳定顺序就写 ORDER BY，不能依赖数据库碰巧按插入顺序返回。

[更多常用写法](../../常用写法速查.md)

## 今天从这里动手

先用一条记录试遍历，再测试三条和零条。

只在遇到时查这些名字：QueryContext、rows.Next、rows.Scan、rows.Err、ORDER BY。

做完后，用自己的话回答：**为什么 Next 循环结束后还需要看 rows.Err？**

[相关文档](https://go.dev/doc/database/querying) · [学习路线](../../README.md) · [上一题](../13-sql-one/README.md) · [下一题](../15-sql-update/README.md)
