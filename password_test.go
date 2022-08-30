package zdpgo_password_generate

import "testing"

func TestGenerateV2(t *testing.T) {
	s := "admin"
	passwordStr := GenerateV2(s)
	if !CheckV2(s, passwordStr) {
		t.Error("密码不正确")
	}
	if CheckV2(s, "admin") {
		t.Error("密码不正确")
	}
}
