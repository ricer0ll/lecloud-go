package lecloud

import "testing"

func TestCredentials(t *testing.T) {
	email := "abc@123.com"
	password := "123"
	accountID := "abc123"

	credentials := NewCredentials(email, password, accountID)

	if credentials.Username != email {
		t.Errorf("username does not match")
	}
	if credentials.Password != password {
		t.Errorf("passwords do not match")
	}
	if credentials.AccountID != accountID {
		t.Errorf("account id do not match")
	}
}
