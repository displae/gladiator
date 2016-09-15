package gladiator

import "testing"

// Validation test for nickname. Nickname should only contains alphanumeric
func TestNicknameValidation(t *testing.T) {
	inputs := []struct {
		Nickname string
		Expected bool
	}{
		{Nickname: "Will Smith", Expected: true},
		{Nickname: "Smith", Expected: true},
		{Nickname: "newbie123", Expected: true},
		{Nickname: "#Walter", Expected: false},
	}
}
