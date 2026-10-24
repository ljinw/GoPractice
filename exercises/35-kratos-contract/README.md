# 35 · 用 Kratos 声明第一个 HTTP 与 gRPC 合同

阶段：Kratos。前置：第 34 题。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 业务为什么需要它

多个服务需要清晰、可生成代码的通信约定。

## 这次做成什么

以健康检查为切口，掌握 protobuf 到 Kratos 处理器的流程。

## 具体需求

采用官方当前 v3 文档，依赖、CLI、代码生成器分别锁定版本。定义 Health.Check，生成 HTTP /health 与 gRPC 入口，均返回 status=ok。先使用独立练习目录和本机端口，不接数据库。生成文件不手改；在自己的 service 实现中写业务。

## 输入与结果

```text
HTTP GET /health → {"status":"ok"}；gRPC Check → status=ok
```

## 怎样算完成

- [ ] HTTP 与 gRPC 两个客户端得到同一含义。
- [ ] 从 proto 重新生成后，手写逻辑仍在。
- [ ] 占用监听端口时启动失败，健康检查不能误报可用。

## 本题常用写法（遇到再看）

【惯例】api/.../v1 放合同，service 做协议转换，biz 放业务，data 放存储；是项目组织约定。【库区别】Kratos 不是 Gin 的别名。

## 今天从这里动手

先只读官方 quick start 的工具要求并记录版本，再创建一个只有 Health 的合同。

先读懂当前输入、输出和失败结果，再写最小实现。框架与依赖到本题才准备；开始前先读 [进阶环境说明](../../进阶环境说明.md)。

[相关文档](https://go-kratos.dev/docs/getting-started/start/) · [业务主线](../../业务主线.md) · [学习路线](../../README.md) · [上一题](../34-business-notification/README.md) · [下一题](../36-kratos-service-split/README.md)
