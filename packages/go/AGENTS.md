# AGENTS（quanttide-secret-toolkit Go 基础库）

面向在 Go 包内工作的编码 agent 的指令。**动手前先读「关键文件」一节**；契约相关改动严格遵守「harness 工作纪律」。路线与交接状态见 [ROADMAP.md](docs/dev-guide/roadmap.md) 与 [STATUS.md](STATUS.md)。

## 本包是什么

标准库的 Go 实现（`pkg/`：envelope/credential/redact）+ 契约测试 runner（`contracttest/`）+ 设计文档（`docs/`）。fixtures 在标准库根 `tests/fixtures/`（本包相对路径上溯三级）。

## 关键文件（按优先级阅读）

| 文件 | 作用 | 何时必读 |
|------|------|----------|
| `README.md` | 布局表、使用示例、开发命令 | 每次工作前 |
| `ROADMAP.md` | 任务状态（S1–S3、H1）与待定决策 | 每次规划前 |
| `STATUS.md` | 交接状态、**待办优先事项**（provider 改引用、credential/redact fixtures、未推送） | 接手任务前 |
| `CONTRIBUTING.md` | 包纪律、契约测试纪律、文档同步清单 | 修改代码前 |
| `docs/dev-guide/*.md` | 设计决策记录（envelope D1 / credential D1 / redact D1） | **新增/变更契约前必读** |
| `docs/user-guide/envelope.md` | 使用契约（信封 JSON 形状与校验规则） | 涉及信封序列化 |
| `contracttest/contract_test.go` | Go runner：断言实现与 fixtures 一致 | 改 runner 前 |
| `../../tests/README.md` | 契约测试架构说明 | 改 fixtures/runner 前 |
| `../../tests/fixtures/envelope.json` | **契约唯一权威**（envelope） | 涉及契约语义时 |
| `Makefile` | `make test/vet/fmt/lint` | 验证时 |

## harness 工作纪律（最高优先级，不可妥协）

harness = `tests/fixtures/`（JSON 契约向量） + `contracttest`（Go runner）。

1. **fixtures 先行**：新增/变更契约，先改 `tests/fixtures/*.json` 与 runner 用例，再改实现；实现与 fixtures 不一致 = 失败
2. **禁止端侧发明语义**：不得在本包或任一端自行新增状态/字段含义——契约演进由 fixtures 驱动各端同步
3. **未知码必须报错**：凭证类型/状态解析遇未知值返回错误，**不用 UNKNOWN 兜底**
4. **零知识红线**：`pkg/envelope` 只做外层结构与格式校验，绝不解析 encrypted 内部内容
5. **主干边界**：纯逻辑（值对象/枚举/纯函数）进 `pkg/`；实体、存储（OSS）、鉴权（JWT）留在 `apps/qtcloud-secret/src/provider`；本包无存储依赖
6. **新增契约四同步**：包文档注释（含 JSON 形状）、单元测试、fixtures+runner（若跨端）、`docs/dev-guide/` 决策记录；另同步 README 布局表、ROADMAP 状态、STATUS 待办

## 验证

```bash
go test ./... && go vet ./...   # 从模块根运行（含 contracttest）
```

提交前核对：测试/vet 通过、文档已同步、STATUS 待办已消化。
