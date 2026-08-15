package envelope

import (
	"strings"
	"testing"
)

func validEnvelope() Envelope {
	return Envelope{
		ID:   "3b6c8e5a-2f4b-4c3d-9a7e-1f2b3c4d5e6f",
		Name: "GitHub 登录",
		Encrypted: EncryptedPayload{
			Algorithm:  "AES-256-GCM",
			KDF:        "Argon2id",
			KDFSalt:    "c2FsdA==",
			Nonce:      "bm9uY2UxMjM0NTY3OA==",
			Ciphertext: "Y2lwaGVydGV4dA==",
		},
	}
}

func TestValidateValid(t *testing.T) {
	e := validEnvelope()
	if err := e.Validate(); err != nil {
		t.Fatalf("合法信封校验失败: %v", err)
	}
}

func TestValidateID(t *testing.T) {
	for name, id := range map[string]string{
		"非 UUID":  "not-a-uuid",
		"非 v4 版本": "3b6c8e5a-2f4b-1c3d-9a7e-1f2b3c4d5e6f",
		"空 id":    "",
	} {
		e := validEnvelope()
		e.ID = id
		if err := e.Validate(); err == nil {
			t.Errorf("%s: id=%q 应校验失败", name, id)
		}
	}
}

func TestValidateName(t *testing.T) {
	e := validEnvelope()
	e.Name = ""
	if err := e.Validate(); err == nil {
		t.Error("空 name 应校验失败")
	}
	e.Name = strings.Repeat("a", MaxNameLength+1)
	if err := e.Validate(); err == nil {
		t.Error("超长 name 应校验失败")
	}
}

func TestValidateEncrypted(t *testing.T) {
	for name, mutate := range map[string]func(*Envelope){
		"algorithm 为空":        func(e *Envelope) { e.Encrypted.Algorithm = "" },
		"kdf 为空":              func(e *Envelope) { e.Encrypted.KDF = "" },
		"kdfSalt 非 base64":    func(e *Envelope) { e.Encrypted.KDFSalt = "!!!" },
		"nonce 为空":            func(e *Envelope) { e.Encrypted.Nonce = "" },
		"ciphertext 非 base64": func(e *Envelope) { e.Encrypted.Ciphertext = "!!!" },
		"ciphertext 为空":       func(e *Envelope) { e.Encrypted.Ciphertext = "" },
	} {
		e := validEnvelope()
		mutate(&e)
		if err := e.Validate(); err == nil {
			t.Errorf("%s 应校验失败", name)
		}
	}
}

func TestParseEnvelope(t *testing.T) {
	body := []byte(`{"id":"3b6c8e5a-2f4b-4c3d-9a7e-1f2b3c4d5e6f","name":"GitHub 登录","encrypted":{"algorithm":"AES-256-GCM","kdf":"Argon2id","kdfSalt":"c2FsdA==","nonce":"bm9uY2UxMjM0NTY3OA==","ciphertext":"Y2lwaGVydGV4dA=="}}`)
	e, err := ParseEnvelope(body)
	if err != nil {
		t.Fatalf("合法请求体解析失败: %v", err)
	}
	if e.ID != "3b6c8e5a-2f4b-4c3d-9a7e-1f2b3c4d5e6f" || e.Name != "GitHub 登录" {
		t.Errorf("解析结果不符: %+v", e)
	}
	if e.Encrypted.Algorithm != "AES-256-GCM" {
		t.Errorf("Algorithm = %q, want AES-256-GCM", e.Encrypted.Algorithm)
	}
}

func TestParseEnvelopeErrors(t *testing.T) {
	cases := []struct {
		name string
		body []byte
	}{
		{"空请求体", nil},
		{"非法 JSON", []byte("{")},
		{"超限请求体", make([]byte, MaxEnvelopeSize+1)},
		{"id 非法", []byte(`{"id":"x","name":"n","encrypted":{"algorithm":"AES-256-GCM","kdf":"Argon2id","kdfSalt":"c2FsdA==","nonce":"bm9uY2UxMjM0NTY3OA==","ciphertext":"Y2lwaGVydGV4dA=="}}`)},
	}
	for _, c := range cases {
		if _, err := ParseEnvelope(c.body); err == nil {
			t.Errorf("%s 应解析失败", c.name)
		}
	}
}
