# quanttide-secret-toolkit Go ROADMAP

工具库目标：**领域主干从 provider 抽出到工具库，以契约测试（共享测试向量）统一各语言实现，从而统一多端**。

主干边界：会校验信封、会定状态、会脱敏的纯逻辑（值对象/枚举/纯函数）进工具库；实体、存储（OSS）、鉴权（JWT 验签）留在应用壳（`qtcloud-secret provider`）。

## 当前阶段：初始化（S1 / S2 / S3 / H1 已完成 2026-08-16）

| # | 优先级 | 任务 | 落点 | 状态 |
|---|--------|------|------|------|
| S1 | P0 | `pkg/envelope` 密文信封契约：外层结构校验（UUID v4 / name 1-256 / algorithm、kdf 必填 / base64 且解码非空 / 信封 ≤64KB），**零知识**：不接触 encrypted 内容 | `pkg/envelope` + 测试；provider `internal/model/envelope.go` 待改引用 | ✅ 已完成 |
| S2 | P0 | `pkg/credential` 凭证类型与状态契约：`Kind`（password/api_key/token）、`Status`（active/rotating/revoked/expired）与迁移表（`CanTransition`），未知值必须报错 | `pkg/credential` + 测试 | ✅ 已完成 |
| S3 | P1 | `pkg/redact` 敏感信息脱敏：`Mask` 通用掩码（保留首尾、中间打码、短值整体打码）+ `MaskToken` 凭证类默认 | `pkg/redact` + 测试 | ✅ 已完成 |
| H1 | P1 | 契约测试骨架：共享测试向量（`tests/fixtures/envelope.json`）+ Go runner（`contracttest/`）——fixtures 是契约唯一权威 | 工具库根 `tests/` + `contracttest/` | ✅ 已完成 |

## 待办（下一步）

- **provider 改引用**：`qtcloud-secret` `internal/model/envelope.go` 删除本地校验实现，委托 `pkg/envelope`（消除双份校验逻辑漂移）
- **fixtures 扩展**：`credential`（类型/状态解析与迁移）与 `redact`（脱敏规则）纳入契约测试骨架（H1 扩展）
- 凭证元数据契约（有效期、轮换周期、负责人）形状设计（见待定决策）

## 待定决策

- **凭证元数据契约**：登记/轮换/回收操作所需元数据形状待定——先有状态契约（S2），实体与存储留在应用壳；若多端出现真实消费者再提炼无存储契约
- **多语言对齐**：Dart/Rust 实现是否启动取决于多端需求；契约测试骨架（H1）已预留通用 runner 结构
- **脱敏规则差异化**：按凭证类型（password/api_key/token）差异化掩码规则是否必要（S3 后置）
