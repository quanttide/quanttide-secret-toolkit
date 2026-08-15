// Package credential 定义凭证类型与状态契约（凭证登记、轮换与回收的生命周期）。
//
// 取值均为英文小写字符串，可直接存入数据库、与接口报文流转；
// 未知值解析必须返回错误，不用 UNKNOWN 兜底（见 docs/dev-guide/credential.md D1）。
package credential

import "fmt"

// Kind 凭证类型。
type Kind string

const (
	KindPassword Kind = "password" // 口令
	KindAPIKey   Kind = "api_key"  // API Key
	KindToken    Kind = "token"    // Token（OAuth/JWT 等）
)

// IsValidKind 报告 k 是否为合法凭证类型（存库前防御校验）。
func (k Kind) IsValidKind() bool {
	switch k {
	case KindPassword, KindAPIKey, KindToken:
		return true
	}
	return false
}

// ParseKind 严格解析凭证类型，未知值返回错误。
func ParseKind(s string) (Kind, error) {
	k := Kind(s)
	if !k.IsValidKind() {
		return "", fmt.Errorf("未知凭证类型 %q", s)
	}
	return k, nil
}

// Status 凭证生命周期状态。
type Status string

const (
	StatusActive   Status = "active"   // 正常使用
	StatusRotating Status = "rotating" // 轮换中（新旧并存）
	StatusRevoked  Status = "revoked"  // 已回收/吊销
	StatusExpired  Status = "expired"  // 已过期
)

// IsValidStatus 报告 s 是否为合法凭证状态（存库前防御校验）。
func (s Status) IsValidStatus() bool {
	switch s {
	case StatusActive, StatusRotating, StatusRevoked, StatusExpired:
		return true
	}
	return false
}

// ParseStatus 严格解析凭证状态，未知值返回错误。
func ParseStatus(s string) (Status, error) {
	st := Status(s)
	if !st.IsValidStatus() {
		return "", fmt.Errorf("未知凭证状态 %q", s)
	}
	return st, nil
}

// 状态迁移表：登记(active) → 轮换(rotating) → 回收(revoked)；expired 由有效期驱动。
var transitions = map[Status]map[Status]bool{
	StatusActive: {
		StatusRotating: true, // 发起轮换
		StatusRevoked:  true, // 直接回收
		StatusExpired:  true, // 有效期自然过期
	},
	StatusRotating: {
		StatusActive:  true, // 轮换完成
		StatusRevoked: true, // 轮换失败回收
	},
	StatusExpired: {
		StatusRevoked: true, // 过期后回收
	},
	StatusRevoked: {}, // 终态
}

// CanTransition 报告凭证是否允许从 from 迁移到 to。
func CanTransition(from, to Status) bool {
	return transitions[from][to]
}
