# 11 · 创建自己的设备数据库

阶段：database/sql + SQLite。前置：10；本阶段先使用命令行。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 这题在完整业务中的位置

设备账本：建立可以在程序重启后继续使用的存储。 [查看完整业务](../../业务主线.md)。

## 这次做成什么

在练习目录创建一个 SQLite 文件和 devices 表。

## 具体需求

使用 database/sql + modernc.org/sqlite，文件放 ./data/devices.db。建立 id INTEGER PRIMARY KEY、name TEXT NOT NULL、ip TEXT NOT NULL UNIQUE、enabled INTEGER NOT NULL DEFAULT 1 CHECK(enabled IN (0,1))。表名 devices。使用 CREATE TABLE IF NOT EXISTS；用 PingContext 验证连接并处理目录/文件错误。程序启动不清表。

## 输入与结果

```text
第一次运行 → 创建数据库与表
第二次运行 → 表仍可用，不丢已有数据
```

## 怎样算完成

- [ ] 数据库文件出现在当前题目的 data 目录。
- [ ] 重复初始化不报“表已存在”，已有资料保留。
- [ ] 不可写路径或将目录当文件时明确失败；不能只凭 sql.Open 无错误就输出成功。

环境连接说明见 [数据库准备](../../数据库准备.md)，只阅读本阶段对应段落。

## 本题常用写法（遇到再看）

```go
import _ "modernc.org/sqlite"
```

- 【语法与用途】空白导入用于执行包的初始化；这里让 SQLite 驱动完成注册，不在代码里直接使用包名。
- 【习惯】`db` 是数据库句柄，`ctx` 是 context.Context；接收上下文的函数通常把 ctx 放第一个参数。
- 【库合同】`*sql.DB` 管理连接池，可复用；它不是“一条只能用一次的连接”。Open 之后仍需实际检查连接。
- 【本题选择】`devices`、`devices.db` 和 data 目录是我们定的名字；驱动名 `sqlite` 要与驱动注册的名字匹配。

[更多常用写法](../../常用写法速查.md)

## 今天从这里动手

先打开根目录的《数据库准备》，只完成连接和 Ping；之后再建表。

只在遇到时查这些名字：sql.Open、PingContext、ExecContext、CREATE TABLE。

做完后，用自己的话回答：**为什么 sql.Open 成功后还要检查实际连接？**

[相关文档](https://go.dev/doc/database/open-handle) · [学习路线](../../README.md) · [上一题](../10-net-echo/README.md) · [下一题](../12-sql-insert/README.md)
