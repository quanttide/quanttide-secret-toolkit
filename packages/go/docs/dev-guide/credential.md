# 设计决策：凭证状态契约（credential D1）

## 背景

凭证管理域需要统一的类型与生命周期语义：口令/API Key/Token 的登记、轮换与回收；状态语义散落在各端会导致漂移。

## 决策

1. **凭证类型** `Kind`：`password` / `api_key` / `token`（英文小写字符串，可直接存库、与接口报文流转）
2. **生命周期状态** `Status`：`active`（正常使用）→ `rotating`（轮换中，新旧并存）→ `revoked`（回收/吊销）；`expired`（过期）由有效期驱动，可回收
3. **迁移表**（`CanTransition`）：active → rotating/revoked/expired；rotating → active/revoked；expired → revoked；revoked 为终态
4. **未知值必须报错**：`ParseKind` / `ParseStatus` 对未知值返回错误，不用 UNKNOWN 兜底（与支付标准库纪律一致）
5. **主干边界**：实体与存储留在应用壳（`qtcloud-secret provider`），本包只有契约与纯函数

## 待定

- 凭证元数据契约（有效期、轮换周期、负责人）形状待定——先有状态契约，实体后置
- 状态集/迁移表演进由 fixtures 驱动跨端同步变更，不得各端自行加状态
