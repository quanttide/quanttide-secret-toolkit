# 密文信封使用契约（envelope）

> 与 `qtcloud-secret` 数据模型同源（`apps/qtcloud-secret/docs/dev-guide/model.md`）。

## JSON 形状

每个密钥一个 JSON 信封对象：「明文元数据 + 密文负载」：

```json
{
  "id": "3b6c8e5a-2f4b-4c3d-9a7e-1f2b3c4d5e6f",
  "name": "GitHub 登录",
  "createdAt": "2026-01-01T00:00:00Z",
  "updatedAt": "2026-01-01T00:00:00Z",
  "encrypted": {
    "algorithm": "AES-256-GCM",
    "kdf": "Argon2id",
    "kdfSalt": "base64...",
    "nonce": "base64...",
    "ciphertext": "base64..."
  }
}
```

| 字段 | 说明 |
|------|------|
| `id` | UUID v4，与对象 key 一致 |
| `name` | 明文元数据，客户端列表/搜索用（明文仅名称，不含密码内容） |
| `createdAt` / `updatedAt` | 展示与排序用，不参与并发控制 |
| `encrypted.algorithm` | 加密算法（如 AES-256-GCM） |
| `encrypted.kdf` / `kdfSalt` | 密钥派生参数（如 Argon2id + 随机盐），随密文存储 |
| `encrypted.nonce` | GCM 随机数（12 字节，每次加密重新生成） |
| `encrypted.ciphertext` | 密码内容密文（GCM 认证标签包含在内） |

## 校验规则（`Validate` / `ParseEnvelope`）

| 字段 | 规则 |
|------|------|
| `id` | 必须是 UUID v4 |
| `name` | 长度 1-256 |
| `encrypted.algorithm` | 非空 |
| `encrypted.kdf` | 非空 |
| `encrypted.kdfSalt` / `nonce` / `ciphertext` | 非空且为 base64；`ciphertext` base64 解码后非空 |
| 请求体（`ParseEnvelope`） | 大小 1-64 KB |

服务端只做外层结构与格式校验，**绝不解析 encrypted 内部内容**（零知识红线，见 `docs/dev-guide/envelope.md`）。
