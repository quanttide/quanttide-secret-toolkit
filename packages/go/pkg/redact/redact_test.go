package redact

import "testing"

func TestMask(t *testing.T) {
	cases := []struct {
		name string
		s    string
		head int
		tail int
		want string
	}{
		{"空串", "", 2, 2, ""},
		{"长度不足整体打码", "abc", 4, 4, "****"},
		{"刚好等于首尾长度", "12345678", 4, 4, "****"},
		{"常规掩码", "sk-abcdef123456", 4, 4, "sk-a*******3456"},
		{"负数按 0 处理", "abcdef", 2, -1, "ab****"},
		{"head 为 0", "abcdef", 0, 2, "****ef"},
		{"中文按字符掩码", "阿里云密钥abcdef", 2, 2, "阿里*******ef"},
	}
	for _, c := range cases {
		if got := Mask(c.s, c.head, c.tail); got != c.want {
			t.Errorf("Mask(%q, %d, %d) = %q, want %q", c.s, c.head, c.tail, got, c.want)
		}
	}
}

func TestMaskToken(t *testing.T) {
	if got := MaskToken("sk-abcdef123456"); got != "sk-a*******3456" {
		t.Errorf("MaskToken = %q, want %q", got, "sk-a*******3456")
	}
	if got := MaskToken("ab"); got != "****" {
		t.Errorf("MaskToken 短值 = %q, want ****", got)
	}
}
