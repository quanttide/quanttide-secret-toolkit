// Package redact 提供敏感信息脱敏（掩码）纯函数，用于日志与审计场景防泄露。
//
// 规则：保留首尾可见字符、中间以 * 打码；长度不足以保留首尾时整体打码，
// 不泄露前缀与长度信息（见 docs/dev-guide/redact.md D1）。
package redact

import "strings"

// Mask 将 s 的中间部分替换为 *，保留前 head 与后 tail 个字符。
// 负数按 0 处理；若 s 不足以同时保留 head+tail 个字符，整体替换为 ****。
func Mask(s string, head, tail int) string {
	if s == "" {
		return ""
	}
	if head < 0 {
		head = 0
	}
	if tail < 0 {
		tail = 0
	}
	r := []rune(s)
	if len(r) <= head+tail {
		return "****"
	}
	return string(r[:head]) + strings.Repeat("*", len(r)-head-tail) + string(r[len(r)-tail:])
}

// MaskToken 凭证类值（API Key/Token/口令）的默认掩码：保留首 4 尾 4，中间打码。
func MaskToken(s string) string {
	return Mask(s, 4, 4)
}
