// Package contracttest 契约测试 Go runner。
//
// 读取工具库根 tests/fixtures/ 下的共享测试向量，断言 Go 实现与契约一致：
// 同一输入 → 同一输出 / 同一拒绝行为。其他语言实现（如 packages/dart、packages/rust）
// 消费同一 fixtures，即完成多端对齐——fixtures 是契约的唯一权威。
package contracttest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/quanttide/quanttide-secret-toolkit/packages/go/pkg/envelope"
)

// fixture 读取共享测试向量。fixtures 位于工具库根 tests/fixtures
// （本包位于 packages/go/contracttest，相对路径上溯三级）。
func fixture(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "tests", "fixtures", name))
	if err != nil {
		t.Fatalf("读取 fixture %s: %v", name, err)
	}
	return b
}

func TestEnvelopeContract(t *testing.T) {
	var fx struct {
		Valid []struct {
			Name string `json:"name"`
			JSON string `json:"json"`
		} `json:"valid"`
		Invalid []struct {
			JSON   string `json:"json"`
			Reason string `json:"reason"`
		} `json:"invalid"`
	}
	b := fixture(t, "envelope.json")
	if err := json.Unmarshal(b, &fx); err != nil {
		t.Fatalf("解析 fixture: %v", err)
	}

	for _, c := range fx.Valid {
		c := c
		t.Run("valid/"+c.Name, func(t *testing.T) {
			if _, err := envelope.ParseEnvelope([]byte(c.JSON)); err != nil {
				t.Fatalf("ParseEnvelope(%s) 错误: %v", c.Name, err)
			}
		})
	}
	for _, c := range fx.Invalid {
		c := c
		t.Run("invalid/"+c.Reason, func(t *testing.T) {
			if _, err := envelope.ParseEnvelope([]byte(c.JSON)); err == nil {
				t.Fatalf("ParseEnvelope 应报错, 期望原因: %s", c.Reason)
			}
		})
	}
}
