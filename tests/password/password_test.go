package main

import (
	"fmt"
	"testing"

	"../../helpers/password"
)

func TestPassword(t *testing.T) {
	hashedPassword := password.GenerateHashedPassword([]byte("MoneyIsTheRealGod"))
	if !password.VerifyPassword([]byte(hashedPassword), []byte("MoneyIsTheRealGod")) {
		t.Error("Password Verification done")
	} else {
		fmt.Println("Test Passed")
	}
}
