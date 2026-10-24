# 36 · 拆成设备服务和巡检服务

阶段：微服务边界。前置：第 35 题。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 业务为什么需要它

设备资料和巡检任务有不同的职责与数据归属。

## 这次做成什么

巡检服务通过 gRPC 获取设备快照，不直接读设备库。

## 具体需求

inventory 服务拥有 devices；inspection 服务拥有 jobs/facts/reports/outbox。各自独立数据库和进程；先用静态本机地址。迁移既有规则时保持一套业务实现，通过适配层连接 Kratos HTTP/gRPC。HTTP 原有 JSON 合同继续测试；protobuf 的 int64 JSON 表示差异需在边界显式转换，不能悄悄改掉外部字段类型。

## 输入与结果

```text
提交巡检 → inspection 调用 inventory.GetDevice → 冻结设备快照 → queued
```

## 怎样算完成

- [ ] 两进程分别启动，任务能取得设备快照。
- [ ] inspection 没有 inventory 库路径或仓储直连。
- [ ] 旧 HTTP 行为测试仍通过；服务停止后明确失败，不伪造设备数据。

## 本题常用写法（遇到再看）

【常见名字】DTO、RPC、client、repository。【边界】拆进程不等于可以共用一个数据库随意跨表读写。

## 今天从这里动手

先把 GetDevice 变成一个 gRPC 调用，再逐步搬出其他设备操作。

先读懂当前输入、输出和失败结果，再写最小实现。框架与依赖到本题才准备；开始前先读 [进阶环境说明](../../进阶环境说明.md)。

[相关文档](https://go-kratos.dev/docs/component/transport/grpc/) · [业务主线](../../业务主线.md) · [学习路线](../../README.md) · [上一题](../35-kratos-contract/README.md) · [下一题](../37-kratos-call-boundary/README.md)
