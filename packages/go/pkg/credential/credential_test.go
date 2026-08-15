package credential

import "testing"

func TestParseKind(t *testing.T) {
	valid := map[string]Kind{
		"password": KindPassword,
		"api_key":  KindAPIKey,
		"token":    KindToken,
	}
	for s, want := range valid {
		k, err := ParseKind(s)
		if err != nil {
			t.Errorf("ParseKind(%q) 错误: %v", s, err)
		}
		if k != want {
			t.Errorf("ParseKind(%q) = %q, want %q", s, k, want)
		}
	}
	for _, s := range []string{"", "PASSWORD", "unknown"} {
		if _, err := ParseKind(s); err == nil {
			t.Errorf("ParseKind(%q) 应报错", s)
		}
	}
}

func TestParseStatus(t *testing.T) {
	valid := map[string]Status{
		"active":   StatusActive,
		"rotating": StatusRotating,
		"revoked":  StatusRevoked,
		"expired":  StatusExpired,
	}
	for s, want := range valid {
		st, err := ParseStatus(s)
		if err != nil {
			t.Errorf("ParseStatus(%q) 错误: %v", s, err)
		}
		if st != want {
			t.Errorf("ParseStatus(%q) = %q, want %q", s, st, want)
		}
	}
	for _, s := range []string{"", "ACTIVE", "unknown"} {
		if _, err := ParseStatus(s); err == nil {
			t.Errorf("ParseStatus(%q) 应报错", s)
		}
	}
}

func TestIsValid(t *testing.T) {
	for _, k := range []Kind{KindPassword, KindAPIKey, KindToken} {
		if !k.IsValidKind() {
			t.Errorf("Kind(%q) 应为合法类型", k)
		}
	}
	if Kind("unknown").IsValidKind() {
		t.Error("未知 Kind 不应合法")
	}
	for _, s := range []Status{StatusActive, StatusRotating, StatusRevoked, StatusExpired} {
		if !s.IsValidStatus() {
			t.Errorf("Status(%q) 应为合法状态", s)
		}
	}
	if Status("unknown").IsValidStatus() {
		t.Error("未知 Status 不应合法")
	}
}

func TestCanTransition(t *testing.T) {
	allowed := [][2]Status{
		{StatusActive, StatusRotating},
		{StatusActive, StatusRevoked},
		{StatusActive, StatusExpired},
		{StatusRotating, StatusActive},
		{StatusRotating, StatusRevoked},
		{StatusExpired, StatusRevoked},
	}
	for _, tt := range allowed {
		if !CanTransition(tt[0], tt[1]) {
			t.Errorf("CanTransition(%s, %s) 应为 true", tt[0], tt[1])
		}
	}
	forbidden := [][2]Status{
		{StatusRevoked, StatusActive},
		{StatusRevoked, StatusExpired},
		{StatusExpired, StatusActive},
		{StatusExpired, StatusRotating},
		{StatusRotating, StatusExpired},
	}
	for _, tt := range forbidden {
		if CanTransition(tt[0], tt[1]) {
			t.Errorf("CanTransition(%s, %s) 应为 false", tt[0], tt[1])
		}
	}
}
