# 契约测试

共享测试向量（fixtures）是各语言实现必须满足的**契约唯一权威**：同一输入 → 同一输出 / 同一拒绝行为。各语言消费同一 fixtures 并断言一致，即完成多端对齐。

## fixtures

| 文件 | 覆盖契约 |
|------|----------|
| `fixtures/envelope.json` | 密文信封外层结构校验（合法信封可解析；非法 id/name/加密字段/JSON 必须拒绝） |

## runner

- Go：`packages/go/contracttest`——`go test ./contracttest`（从 Go 模块根运行）
- 新增语言（如 `packages/dart`、`packages/rust`）：实现同一契约，消费同一 fixtures 断言输出一致

## 待补

- `credential`（凭证类型/状态解析与迁移）与 `redact`（脱敏规则）fixtures 纳入骨架（见 `packages/go/ROADMAP.md` 待办）
