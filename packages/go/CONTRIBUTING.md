# 贡献指南（quanttide-secret-toolkit Go 基础库）

## 本包定位

Go 基础库承载**领域主干**：会校验信封、会定状态、会脱敏的纯逻辑（值对象/枚举/纯函数）。实体、存储（OSS）、鉴权（JWT 验签）留在应用壳（`apps/qtcloud-secret/src/provider`）。

## 包纪律

1. **纯逻辑**：包内不引入 gorm/数据库/http client 等存储与 IO 依赖；当前零外部依赖（纯标准库）
2. **契约即文档**：每个包的文档注释写清契约语义与 JSON 形状；取值用英文小写字符串，可直接存库
3. **防御校验**：枚举/状态提供 `IsValid*` 供存库前校验；未知值解析**必须报错**，不用 UNKNOWN 兜底
4. **零知识红线**：密文信封只做外层结构校验，绝不解析/解密 `encrypted` 内容

## 契约测试纪律（harness）

harness = `tests/fixtures/`（JSON 契约向量，契约唯一权威）+ `contracttest/`（Go runner）：

1. **fixtures 先行**：新增/变更契约，先改 `tests/fixtures/*.json` 与 runner 用例，再改实现；实现与 fixtures 不一致 = 失败
2. 现有 fixtures：`envelope.json`；**credential / redact 待补**（ROADMAP 待办）
3. 禁止各端自行新增状态/字段含义——契约演进由 fixtures 驱动两端同步

## 文档同步（新增/变更契约时缺一不可）

| 项 | 要求 |
|----|------|
| 包文档注释 | 含契约语义与 JSON 形状 |
| 单元测试 | 与包同目录 `*_test.go` |
| fixtures + runner | 若契约跨端（信封/凭证/脱敏） |
| `docs/dev-guide/*.md` | 设计决策记录（新增契约前先读已有 D1） |
| `README.md` | 布局表同步 |
| `ROADMAP.md` / `STATUS.md` | 任务状态与交接待办同步 |

## 开发命令

```bash
make test    # go test ./...（含 contracttest，从模块根运行）
make vet     # go vet ./...
make fmt     # gofmt -w .
make lint    # golangci-lint（需先安装）
```

契约测试单独运行：`go test ./contracttest`（消费 `../../tests/fixtures/`）。

## 提交规范

Conventional Commits，中文描述，作用域 `toolkit`：

```
feat(toolkit): pkg/envelope 密文信封契约（从 provider model 抽主干）
docs(toolkit): ROADMAP 对齐当前进度
```

提交前核对：`go test ./... && go vet ./...` 通过、STATUS.md 待办已消化、文档已同步。
