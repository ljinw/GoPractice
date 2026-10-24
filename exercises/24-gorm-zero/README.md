# 24 · 把启用状态真的改成 false

阶段：gorm。前置：23。

## 先学再练：本题串讲

打开本目录的 [知识串讲](知识串讲.md)，先学这题需要的语法、原理和常用写法，再回来实现下面的业务需求。

配套：[多轮练习册：变式、排错、源码与C#](练习册.md)。先完成当前小步骤，再按练习册安排重复与复习。

## 这题在完整业务中的位置

维护开关：为第 25 题“停用设备不接新巡检”准备准确的存储行为。 [查看完整业务](../../业务主线.md)。

## 这次做成什么

用 GORM 正确更新一个布尔值为 false，避免把零值更新漏掉。

## 具体需求

先做存储函数 SetEnabled(id, enabled)，不新增 HTTP 接口。只修改指定 ID 的 enabled，false 对应 SQLite 的 0。缺记录返回 not found；重复设置同一值可成功。选择明确包含该字段的更新方式；不要用可能创建新记录的 Save 来代替更新。

## 输入与结果

```text
enabled=true → 设置 false → 重新查询得到 false
再设 false → 仍为 false，行数不变
```

## 怎样算完成

- [ ] true→false、false→true 两方向都正确。
- [ ] 设置 false 后重启并重新查询仍是 false。
- [ ] 另一台设备不变；不存在 ID 不新增记录；数据库错误正常返回。

## 本题常用写法（遇到再看）

```go
changes := map[string]any{"enabled": false}
```

- 【语法】any 是 interface{} 的别名；这里表示 map 的值允许不同类型。当前只放一个 bool。
- 【库合同】GORM 传统 API 的 Updates(struct) 默认忽略零值；false、0、空字符串都可能受影响。
- 【常见用法】明确更新单列，或用 map/Select 明确包含字段，表达“我就是要写入 false”。
- 【业务合同】先限定目标 ID，检查 Error 和影响范围；Update 与 Save 的语义不同，保存接口的名字不能代替查阅实际行为。

[更多常用写法](../../常用写法速查.md)

## 今天从这里动手

先写“true 改成 false 后必须读到 false”的失败测试，再选择更新写法。

只在遇到时查这些名字：Update、Updates、map、Select、RowsAffected。

做完后，用自己的话回答：**为什么更新时 false 有时会被当成“未提供这个字段”？**

[相关文档](https://gorm.io/docs/update.html) · [学习路线](../../README.md) · [上一题](../23-gorm-store/README.md) · [下一题](../25-business-device-enabled/README.md)
