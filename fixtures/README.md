# 本地业务样例

- [inspection-samples.json](inspection-samples.json)：正常、异常、缺失三个合成采样。
- [knowledge.json](knowledge.json)：三条自编教学知识，供 Retriever/RAG 练习。

使用时按 device_name/fixture_key 建立本地 Collector 映射，不向这些示例 IP 发起连接。expected_health 是测试的期望值，采集器必须根据 value 和规则计算，不直接复制它作为判断实现。

fact_key 是样例内的标签；真正执行时按 job_id + fact_key 生成该任务唯一 fact_id，避免不同任务引用串用。知识引用使用 kb_id + version。注入未知事实、坏 JSON、文件损坏等失败场景时，复制到测试临时目录后修改，保留原始样例。
