// Command secret 演示 quanttide-secret-toolkit Go 基础库的用法。
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/quanttide/quanttide-secret-toolkit/packages/go/pkg/credential"
	"github.com/quanttide/quanttide-secret-toolkit/packages/go/pkg/envelope"
	"github.com/quanttide/quanttide-secret-toolkit/packages/go/pkg/redact"
)

func main() {
	// 信封：外层结构校验（零知识：不解密、不接触密文内容）
	env := &envelope.Envelope{
		ID:   "3b6c8e5a-2f4b-4c3d-9a7e-1f2b3c4d5e6f",
		Name: "GitHub 登录",
		Encrypted: envelope.EncryptedPayload{
			Algorithm:  "AES-256-GCM",
			KDF:        "Argon2id",
			KDFSalt:    "c2FsdA==",
			Nonce:      "bm9uY2UxMjM0NTY3OA==",
			Ciphertext: "Y2lwaGVydGV4dA==",
		},
	}
	if err := env.Validate(); err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(env)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("信封: %s\n", b)

	// 凭证：类型与状态契约，取值可直接存库、与接口报文流转
	k, err := credential.ParseKind("api_key")
	if err != nil {
		log.Fatal(err)
	}
	s := credential.StatusActive
	fmt.Printf("凭证类型: %s, 状态: %s, 可发起轮换: %v\n",
		k, s, credential.CanTransition(s, credential.StatusRotating))

	// 脱敏：日志/审计不泄露完整凭证
	fmt.Printf("脱敏: %s\n", redact.MaskToken("sk-abcdef123456"))
}
