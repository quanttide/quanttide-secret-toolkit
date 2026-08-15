# quanttide-secret-toolkit Go 基础库

量潮机密管理工具箱的 Go 基础库，提供机密管理领域的通用基础能力：密文信封契约、凭证生命周期状态、敏感信息脱敏。遵循 Go 社区标准布局：`pkg/` 为可被外部导入的公共包，`internal/` 为库内部实现，`examples/` 为使用示例。

## 布局

| 路径 | 说明 |
|------|------|
| `pkg/envelope` | 密文信封契约（值对象 + 外层结构校验）：UUID v4 / name 1-256 / algorithm、kdf 必填 / base64 字段 / 信封 ≤64 KB；**零知识**：绝不解析 encrypted 内容 |
| `pkg/credential` | 凭证类型与状态契约：`Kind`（password/api_key/token）、`Status`（active/rotating/revoked/expired）与迁移校验（`CanTransition`） |
| `pkg/redact` | 敏感信息脱敏：`Mask` 通用掩码（保留首尾、中间打码）与 `MaskToken` 凭证类默认掩码 |
| `internal/` | 库的内部实现，外部不可导入（预留） |
| `contracttest/` | 契约测试 Go runner（消费工具库根 `tests/fixtures/`，fixtures 是契约唯一权威） |
| `docs/` | 使用指南与设计决策记录（`user-guide/envelope.md` 信封 JSON 形状；`dev-guide/envelope.md`、`dev-guide/credential.md`、`dev-guide/redact.md`） |
| `examples/` | 使用示例（`go run ./examples/secret`） |

## 使用

```go
import (
	"github.com/quanttide/quanttide-secret-toolkit/packages/go/pkg/credential"
	"github.com/quanttide/quanttide-secret-toolkit/packages/go/pkg/envelope"
	"github.com/quanttide/quanttide-secret-toolkit/packages/go/pkg/redact"
)

// 信封：外层结构校验（不解密、不接触内容）
env := &envelope.Envelope{ID: "...", Name: "GitHub 登录", Encrypted: envelope.EncryptedPayload{...}}
err := env.Validate()

// 凭证状态：取值可直接存入数据库、与接口报文流转
s := credential.StatusRotating
// s == "rotating"

// 脱敏：日志/审计不泄露完整凭证
redact.MaskToken("sk-abcdef123456") // "sk-a*******3456"
```

## 契约测试

共享测试向量（JSON fixtures，工具库根 `tests/fixtures/`）是契约**唯一权威**，各语言实现消费同一份断言一致。Go runner 位于 `contracttest/`：

```bash
go test ./contracttest
```

## 文档

| 文档 | 内容 |
|------|------|
| [ROADMAP.md](ROADMAP.md) | 任务状态与待定决策（S1–S3、H1） |
| [STATUS.md](STATUS.md) | 交接状态与待办优先事项 |
| [CONTRIBUTING.md](CONTRIBUTING.md) | 包纪律、契约测试纪律、文档同步清单 |
| [AGENTS.md](AGENTS.md) | agent 工作索引与 harness 纪律 |
| `docs/dev-guide/` | 设计决策记录（改契约前先读） |
| `docs/user-guide/` | 使用契约 |

## 开发

```bash
make test    # 单元测试
make vet     # 静态检查
make fmt     # 格式化
make lint    # golangci-lint（需先安装）
```
