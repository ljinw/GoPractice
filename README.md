# Go 项目练习：每个练习点配自己的知识串讲

一个持续演进的项目：**设备巡检与 AI 报告助手**。目标是既能写业务，也能部署、监控、排障和恢复。

完整阅读：[Go开发与运维知识串讲](Go开发与运维知识串讲.md)。按项目练习：下方目录。两种方式同时保留。

新的学习安排：[多轮练习与岗位能力](学习方式与岗位能力路线.md) · [Go/C#及源码对照](Go与CSharp源码对照.md) · [安全学习与12个防御实验](安全学习与项目练习.md)。69个练习点各有9轮练习，共621个练习检查项；它们不替代原来的207项业务验收，也不代表已经完成或只练这些就足够。

当前建议按每周16–18小时主动学习安排：公司阅读，晚上与周末动手。完整目录是长期练习库，近期先推进01–19的基础主线；具体节奏和阶段目标见[个人学习安排](学习方式与岗位能力路线.md)。

第一阶段先按[标准库学习路线](第一阶段标准库学习路线.md)学fmt/log、HTTP服务端与客户端，再学net地址与TCP；01–10题已补逐API解释与独立小实验。

## 每题怎么学

每个 exercises 子目录把材料放在一起：

- README.md：本题业务问题、目标行为、前置题目、输入输出、步骤和验收。
- 知识串讲.md：本题要复习的语法、新知识、原理、常用写法与失败排查；读完直接接本题实现。
- 练习册.md：本题多轮变式、排错、源码阅读、C#对照与隔题复习。
- 你自己的代码、配置和实验记录：练习时逐步补入；本次整理没有代写业务答案或部署服务。

先打开 [开始这里](开始这里.md)，或 [01题讲义](exercises/01-http-health/知识串讲.md)。每次只走当前练习的一小步，无需先读完整本理论。

## 项目怎样连续演进

01–24练局部能力；25–48组合成巡检项目并完成第一轮本地交付；49–69继续用这套项目补数据库进阶和运维。题号表示学习阶段，前置依赖写在题目中；可按依赖复习或跳转。

业务合同见 [业务主线](业务主线.md)。技术对照题保持相同接口、权限和事实规则；单次运行选择一套权威实现，避免多个适配同时重复产生效果。

## 学习阶段

| 练习点 | 阶段 | 在项目中完成什么 |
|---|---|---|
| 01–10 | HTTP 与网络 | 请求响应、JSON、客户端、测试、IP/CIDR、本机连接 |
| 11–19 | 数据库接到接口 | SQL、事务、设备增查改删与 HTTP 合同 |
| 20–24 | 框架与 ORM | chi、Gin、GORM，保持相同业务行为 |
| 25–34 | 完整业务与中间件 | 设备启停、任务、报告、权限、幂等、缓存、队列与通知 |
| 35–40 | 微服务与环境 | Kratos、RPC、服务发现、追踪与 Compose |
| 41–48 | AI 与第一轮交付 | Eino、证据、RAG、工具、流式、基础闭环验收 |
| 49–54 | 数据库进阶 | 并发更新、锁、死锁、优化、恢复、MongoDB |
| 55–56 | 消息系统对照 | RabbitMQ 通知交付、Kafka 事实回放 |
| 57–62 | 单机运维与可观测性 | Linux、Nginx、日志、监控、异步追踪与容器数据 |
| 63–64 | Kubernetes | 部署、探针、资源、可用性与排障 |
| 65–69 | 上线迭代与最终交接 | 流水线、兼容迁移、灰度验证、回滚与故障演练 |

## 逐题打开：讲义与需求在同一个目录

共 **69个练习点**：原48题保留业务需求与验收，新增21个进阶练习点。

| 题号 | 项目中的练习点 | 本题知识 | 动手要求 | 反复练习 |
|---|---|---|---|---|
| 01 | 让浏览器收到一句 ok | [知识串讲](exercises/01-http-health/知识串讲.md) | [需求与验收](exercises/01-http-health/README.md) | [9轮练习](exercises/01-http-health/练习册.md) |
| 02 | 从网址拿到设备名称 | [知识串讲](exercises/02-http-query/知识串讲.md) | [需求与验收](exercises/02-http-query/README.md) | [9轮练习](exercises/02-http-query/练习册.md) |
| 03 | 把设备信息返回为 JSON | [知识串讲](exercises/03-http-json/知识串讲.md) | [需求与验收](exercises/03-http-json/README.md) | [9轮练习](exercises/03-http-json/练习册.md) |
| 04 | 接收一份设备资料 | [知识串讲](exercises/04-http-post/知识串讲.md) | [需求与验收](exercises/04-http-post/README.md) | [9轮练习](exercises/04-http-post/练习册.md) |
| 05 | 用 Go 访问自己的 HTTP 服务 | [知识串讲](exercises/05-http-client/知识串讲.md) | [需求与验收](exercises/05-http-client/README.md) | [9轮练习](exercises/05-http-client/练习册.md) |
| 06 | 让程序替你检查接口 | [知识串讲](exercises/06-http-tests/知识串讲.md) | [需求与验收](exercises/06-http-tests/README.md) | [9轮练习](exercises/06-http-tests/练习册.md) |
| 07 | 判断一个 IP 写得是否合法 | [知识串讲](exercises/07-net-ip/知识串讲.md) | [需求与验收](exercises/07-net-ip/README.md) | [9轮练习](exercises/07-net-ip/练习册.md) |
| 08 | 判断 IP 是否属于某个网段 | [知识串讲](exercises/08-net-cidr/知识串讲.md) | [需求与验收](exercises/08-net-cidr/README.md) | [9轮练习](exercises/08-net-cidr/练习册.md) |
| 09 | 检查本机一个 TCP 端口 | [知识串讲](exercises/09-net-connect/知识串讲.md) | [需求与验收](exercises/09-net-connect/README.md) | [9轮练习](exercises/09-net-connect/练习册.md) |
| 10 | 让本机 TCP 服务回复一行文字 | [知识串讲](exercises/10-net-echo/知识串讲.md) | [需求与验收](exercises/10-net-echo/README.md) | [9轮练习](exercises/10-net-echo/练习册.md) |
| 11 | 创建自己的设备数据库 | [知识串讲](exercises/11-sql-open/知识串讲.md) | [需求与验收](exercises/11-sql-open/README.md) | [9轮练习](exercises/11-sql-open/练习册.md) |
| 12 | 保存一台设备 | [知识串讲](exercises/12-sql-insert/知识串讲.md) | [需求与验收](exercises/12-sql-insert/README.md) | [9轮练习](exercises/12-sql-insert/练习册.md) |
| 13 | 按编号找一台设备 | [知识串讲](exercises/13-sql-one/知识串讲.md) | [需求与验收](exercises/13-sql-one/README.md) | [9轮练习](exercises/13-sql-one/练习册.md) |
| 14 | 列出全部设备 | [知识串讲](exercises/14-sql-list/知识串讲.md) | [需求与验收](exercises/14-sql-list/README.md) | [9轮练习](exercises/14-sql-list/练习册.md) |
| 15 | 修改一台设备的名称 | [知识串讲](exercises/15-sql-update/知识串讲.md) | [需求与验收](exercises/15-sql-update/README.md) | [9轮练习](exercises/15-sql-update/练习册.md) |
| 16 | 删除一台假设备 | [知识串讲](exercises/16-sql-delete/知识串讲.md) | [需求与验收](exercises/16-sql-delete/README.md) | [9轮练习](exercises/16-sql-delete/练习册.md) |
| 17 | 让两条新增一起成功或一起失败 | [知识串讲](exercises/17-sql-transaction/知识串讲.md) | [需求与验收](exercises/17-sql-transaction/README.md) | [9轮练习](exercises/17-sql-transaction/练习册.md) |
| 18 | 通过 HTTP 查询已保存的设备 | [知识串讲](exercises/18-http-sql-read/知识串讲.md) | [需求与验收](exercises/18-http-sql-read/README.md) | [9轮练习](exercises/18-http-sql-read/练习册.md) |
| 19 | 通过 HTTP 保存设备 | [知识串讲](exercises/19-http-sql-create/知识串讲.md) | [需求与验收](exercises/19-http-sql-create/README.md) | [9轮练习](exercises/19-http-sql-create/练习册.md) |
| 20 | 用 chi 接管已有路由 | [知识串讲](exercises/20-chi-router/知识串讲.md) | [需求与验收](exercises/20-chi-router/README.md) | [9轮练习](exercises/20-chi-router/练习册.md) |
| 21 | 给每次请求带一个编号 | [知识串讲](exercises/21-chi-request-id/知识串讲.md) | [需求与验收](exercises/21-chi-request-id/README.md) | [9轮练习](exercises/21-chi-request-id/练习册.md) |
| 22 | 用 gin 写同一组接口 | [知识串讲](exercises/22-gin-router/知识串讲.md) | [需求与验收](exercises/22-gin-router/README.md) | [9轮练习](exercises/22-gin-router/练习册.md) |
| 23 | 用 GORM 保存和读取设备 | [知识串讲](exercises/23-gorm-store/知识串讲.md) | [需求与验收](exercises/23-gorm-store/README.md) | [9轮练习](exercises/23-gorm-store/练习册.md) |
| 24 | 把启用状态真的改成 false | [知识串讲](exercises/24-gorm-zero/知识串讲.md) | [需求与验收](exercises/24-gorm-zero/README.md) | [9轮练习](exercises/24-gorm-zero/练习册.md) |
| 25 | 让停用设备退出后续巡检 | [知识串讲](exercises/25-business-device-enabled/知识串讲.md) | [需求与验收](exercises/25-business-device-enabled/README.md) | [9轮练习](exercises/25-business-device-enabled/练习册.md) |
| 26 | 提交一次有编号的巡检任务 | [知识串讲](exercises/26-business-create-job/知识串讲.md) | [需求与验收](exercises/26-business-create-job/README.md) | [9轮练习](exercises/26-business-create-job/练习册.md) |
| 27 | 执行巡检并留下可核查的事实 | [知识串讲](exercises/27-business-run-job/知识串讲.md) | [需求与验收](exercises/27-business-run-job/README.md) | [9轮练习](exercises/27-business-run-job/练习册.md) |
| 28 | 给用户一份不会随数据变化的报告 | [知识串讲](exercises/28-business-report/知识串讲.md) | [需求与验收](exercises/28-business-report/README.md) | [9轮练习](exercises/28-business-report/练习册.md) |
| 29 | 只允许用户查看自己的任务 | [知识串讲](exercises/29-middleware-auth/知识串讲.md) | [需求与验收](exercises/29-middleware-auth/README.md) | [9轮练习](exercises/29-middleware-auth/练习册.md) |
| 30 | 重复点击也只创建一个任务 | [知识串讲](exercises/30-middleware-idempotency/知识串讲.md) | [需求与验收](exercises/30-middleware-idempotency/README.md) | [9轮练习](exercises/30-middleware-idempotency/练习册.md) |
| 31 | 缓存已完成报告并保留权限检查 | [知识串讲](exercises/31-redis-report-cache/知识串讲.md) | [需求与验收](exercises/31-redis-report-cache/README.md) | [9轮练习](exercises/31-redis-report-cache/练习册.md) |
| 32 | 把耗时巡检交给后台处理 | [知识串讲](exercises/32-redis-job-queue/知识串讲.md) | [需求与验收](exercises/32-redis-job-queue/README.md) | [9轮练习](exercises/32-redis-job-queue/练习册.md) |
| 33 | 处理 worker 中断与重试 | [知识串讲](exercises/33-redis-worker-recovery/知识串讲.md) | [需求与验收](exercises/33-redis-worker-recovery/README.md) | [9轮练习](exercises/33-redis-worker-recovery/练习册.md) |
| 34 | 把报告就绪消息送到本机接收器 | [知识串讲](exercises/34-business-notification/知识串讲.md) | [需求与验收](exercises/34-business-notification/README.md) | [9轮练习](exercises/34-business-notification/练习册.md) |
| 35 | 用 Kratos 声明第一个 HTTP 与 gRPC 合同 | [知识串讲](exercises/35-kratos-contract/知识串讲.md) | [需求与验收](exercises/35-kratos-contract/README.md) | [9轮练习](exercises/35-kratos-contract/练习册.md) |
| 36 | 拆成设备服务和巡检服务 | [知识串讲](exercises/36-kratos-service-split/知识串讲.md) | [需求与验收](exercises/36-kratos-service-split/README.md) | [9轮练习](exercises/36-kratos-service-split/练习册.md) |
| 37 | 让远程调用有身份与截止时间 | [知识串讲](exercises/37-kratos-call-boundary/知识串讲.md) | [需求与验收](exercises/37-kratos-call-boundary/README.md) | [9轮练习](exercises/37-kratos-call-boundary/练习册.md) |
| 38 | 换端口后仍能找到设备服务 | [知识串讲](exercises/38-kratos-discovery/知识串讲.md) | [需求与验收](exercises/38-kratos-discovery/README.md) | [9轮练习](exercises/38-kratos-discovery/练习册.md) |
| 39 | 沿一次请求找到慢在哪里 | [知识串讲](exercises/39-observability-trace/知识串讲.md) | [需求与验收](exercises/39-observability-trace/README.md) | [9轮练习](exercises/39-observability-trace/练习册.md) |
| 40 | 让本机练习环境可重复启动 | [知识串讲](exercises/40-compose-environment/知识串讲.md) | [需求与验收](exercises/40-compose-environment/README.md) | [9轮练习](exercises/40-compose-environment/练习册.md) |
| 41 | 让 Eino 接住一份巡检报告 | [知识串讲](exercises/41-eino-first-analysis/知识串讲.md) | [需求与验收](exercises/41-eino-first-analysis/README.md) | [9轮练习](exercises/41-eino-first-analysis/练习册.md) |
| 42 | 只保存有证据的 AI 分析 | [知识串讲](exercises/42-eino-structured-output/知识串讲.md) | [需求与验收](exercises/42-eino-structured-output/README.md) | [9轮练习](exercises/42-eino-structured-output/练习册.md) |
| 43 | 把分析过程拆成可检查的步骤 | [知识串讲](exercises/43-eino-graph/知识串讲.md) | [需求与验收](exercises/43-eino-graph/README.md) | [9轮练习](exercises/43-eino-graph/练习册.md) |
| 44 | 让建议引用本地排障知识 | [知识串讲](exercises/44-eino-rag/知识串讲.md) | [需求与验收](exercises/44-eino-rag/README.md) | [9轮练习](exercises/44-eino-rag/练习册.md) |
| 45 | 让 Agent 自己选择只读查询工具 | [知识串讲](exercises/45-eino-agent-tools/知识串讲.md) | [需求与验收](exercises/45-eino-agent-tools/README.md) | [9轮练习](exercises/45-eino-agent-tools/练习册.md) |
| 46 | 让用户看见分析正在进行 | [知识串讲](exercises/46-eino-sse-progress/知识串讲.md) | [需求与验收](exercises/46-eino-sse-progress/README.md) | [9轮练习](exercises/46-eino-sse-progress/练习册.md) |
| 47 | 把假模型替换为一次真实模型调用 | [知识串讲](exercises/47-eino-live-provider/知识串讲.md) | [需求与验收](exercises/47-eino-live-provider/README.md) | [9轮练习](exercises/47-eino-live-provider/练习册.md) |
| 48 | 交付一个能演示和复现的巡检助手 | [知识串讲](exercises/48-business-e2e/知识串讲.md) | [需求与验收](exercises/48-business-e2e/README.md) | [9轮练习](exercises/48-business-e2e/练习册.md) |
| 49 | 两个人改设备时不悄悄覆盖 | [知识串讲](exercises/49-sql-concurrent-device/知识串讲.md) | [需求与验收](exercises/49-sql-concurrent-device/README.md) | [9轮练习](exercises/49-sql-concurrent-device/练习册.md) |
| 50 | 批量维护设备遇到死锁后安全重试 | [知识串讲](exercises/50-sql-deadlock-retry/知识串讲.md) | [需求与验收](exercises/50-sql-deadlock-retry/README.md) | [9轮练习](exercises/50-sql-deadlock-retry/练习册.md) |
| 51 | 让报告列表数据多了也能解释性能 | [知识串讲](exercises/51-sql-report-query-plan/知识串讲.md) | [需求与验收](exercises/51-sql-report-query-plan/README.md) | [9轮练习](exercises/51-sql-report-query-plan/练习册.md) |
| 52 | 把巡检数据恢复到另一套环境 | [知识串讲](exercises/52-database-restore-drill/知识串讲.md) | [需求与验收](exercises/52-database-restore-drill/README.md) | [9轮练习](exercises/52-database-restore-drill/练习册.md) |
| 53 | 用文档模型保存一份不可变报告 | [知识串讲](exercises/53-mongodb-report-document/知识串讲.md) | [需求与验收](exercises/53-mongodb-report-document/README.md) | [9轮练习](exercises/53-mongodb-report-document/练习册.md) |
| 54 | 为文档报告查询建立索引证据 | [知识串讲](exercises/54-mongodb-report-index/知识串讲.md) | [需求与验收](exercises/54-mongodb-report-index/README.md) | [9轮练习](exercises/54-mongodb-report-index/练习册.md) |
| 55 | 用 RabbitMQ 可靠交付报告就绪事件 | [知识串讲](exercises/55-rabbitmq-notification/知识串讲.md) | [需求与验收](exercises/55-rabbitmq-notification/README.md) | [9轮练习](exercises/55-rabbitmq-notification/练习册.md) |
| 56 | 从 Kafka 回放事实重建报告统计 | [知识串讲](exercises/56-kafka-fact-replay/知识串讲.md) | [需求与验收](exercises/56-kafka-fact-replay/README.md) | [9轮练习](exercises/56-kafka-fact-replay/练习册.md) |
| 57 | 把巡检助手作为 Linux 服务可靠启停 | [知识串讲](exercises/57-linux-service-lifecycle/知识串讲.md) | [需求与验收](exercises/57-linux-service-lifecycle/README.md) | [9轮练习](exercises/57-linux-service-lifecycle/练习册.md) |
| 58 | 通过 Nginx 访问巡检接口和分析进度 | [知识串讲](exercises/58-nginx-project-gateway/知识串讲.md) | [需求与验收](exercises/58-nginx-project-gateway/README.md) | [9轮练习](exercises/58-nginx-project-gateway/练习册.md) |
| 59 | 凭任务编号查到一次失败的完整经过 | [知识串讲](exercises/59-logs-job-correlation/知识串讲.md) | [需求与验收](exercises/59-logs-job-correlation/README.md) | [9轮练习](exercises/59-logs-job-correlation/练习册.md) |
| 60 | 让任务积压被监控发现并正确恢复 | [知识串讲](exercises/60-metrics-task-alerts/知识串讲.md) | [需求与验收](exercises/60-metrics-task-alerts/README.md) | [9轮练习](exercises/60-metrics-task-alerts/练习册.md) |
| 61 | 沿一次任务追到队列和 AI 分析 | [知识串讲](exercises/61-trace-async-analysis/知识串讲.md) | [需求与验收](exercises/61-trace-async-analysis/README.md) | [9轮练习](exercises/61-trace-async-analysis/练习册.md) |
| 62 | 重建容器后巡检数据仍能使用 | [知识串讲](exercises/62-container-persistent-project/知识串讲.md) | [需求与验收](exercises/62-container-persistent-project/README.md) | [9轮练习](exercises/62-container-persistent-project/练习册.md) |
| 63 | 在 Kubernetes 维持两份可用 API | [知识串讲](exercises/63-k8s-api-deployment/知识串讲.md) | [需求与验收](exercises/63-k8s-api-deployment/README.md) | [9轮练习](exercises/63-k8s-api-deployment/练习册.md) |
| 64 | 定位 Pod 在运行却不能接请求 | [知识串讲](exercises/64-k8s-readiness-diagnosis/知识串讲.md) | [需求与验收](exercises/64-k8s-readiness-diagnosis/README.md) | [9轮练习](exercises/64-k8s-readiness-diagnosis/练习册.md) |
| 65 | 让一次代码改动产出可追溯的发布包 | [知识串讲](exercises/65-ci-release-artifact/知识串讲.md) | [需求与验收](exercises/65-ci-release-artifact/README.md) | [9轮练习](exercises/65-ci-release-artifact/练习册.md) |
| 66 | 新增报告字段时让旧版仍能运行 | [知识串讲](exercises/66-compatible-report-migration/知识串讲.md) | [需求与验收](exercises/66-compatible-report-migration/README.md) | [9轮练习](exercises/66-compatible-report-migration/练习册.md) |
| 67 | 把新版先交给一组教学用户验证 | [知识串讲](exercises/67-canary-project-release/知识串讲.md) | [需求与验收](exercises/67-canary-project-release/README.md) | [9轮练习](exercises/67-canary-project-release/练习册.md) |
| 68 | 灰度失败后把巡检业务恢复到旧版 | [知识串讲](exercises/68-rollback-project-release/知识串讲.md) | [需求与验收](exercises/68-rollback-project-release/README.md) | [9轮练习](exercises/68-rollback-project-release/练习册.md) |
| 69 | 交付一套能开发、能维护、能恢复的巡检助手 | [知识串讲](exercises/69-devops-project-handover/知识串讲.md) | [需求与验收](exercises/69-devops-project-handover/README.md) | [9轮练习](exercises/69-devops-project-handover/练习册.md) |

## 复习和环境

[语法复习地图](语法复习地图.md) 定位多次复习；[常用写法速查](常用写法速查.md) 区分语法、合同和习惯；[知识导航](知识导航.md) 按主题指向题目。

环境按题准备，见 [进阶环境说明](进阶环境说明.md)。不在第一天安装全部组件，数据库和队列对照分开运行。

## 验证边界

当前交付是课程材料与配套实验要求。教材结构通过检查不代表业务通过验收；数据库、消息、容器、集群需实际本地验证。真实设备、生产、通知和付费模型不由本地练习自动授权。

2026-10-24教材校验：69题各保留需求、讲义与9轮练习；原207项验收逐字保留。Go/C#最小函数对照程序已运行且输出一致。源码阅读任务和部署实验仍由后续学习逐项完成，未替你勾选。
