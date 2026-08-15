# 设计决策：密文信封契约（envelope D1）

## 背景

`qtcloud-secret` provider 的 `internal/model/envelope.go` 已实现信封外层校验（纯逻辑）；客户端（未来多端）与审计侧需要同一份校验语义，双份实现存在漂移风险。

## 决策

1. **信封契约抽入工具库** `pkg/envelope`：纯值对象 + 校验函数，无存储/IO 依赖；provider 改为委托引用（消除双份校验逻辑漂移）
2. **零知识红线**：只做外层结构与格式校验，绝不解析/解密 `encrypted` 内容——那是客户端零知识的领地（对齐 `apps/qtcloud-secret/docs/dev-guide/security.md`）
3. **校验规则即契约**：UUID v4、name 1-256、algorithm/kdf 必填、kdfSalt/nonce/ciphertext 必须 base64 且解码后非空、信封 ≤64 KB
4. **契约演进由 fixtures 驱动**：`tests/fixtures/envelope.json` 是唯一权威，跨端同步变更，不得各端自行加字段/规则

## 影响

- provider `internal/model/envelope.go` 删除本地校验实现，委托 `pkg/envelope`
- 多端（如客户端 Dart）实现同一契约，消费同一 fixtures
