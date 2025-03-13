package encrypt

import (
	"bytes"
	"testing"
)

func TestEncryptionDecryption(t *testing.T) {
	testCases := []struct {
		originalData string
		password     string
	}{
		{"Hello, World!", "simple-password"},
		{"This is a test message that is a bit longer than the first one", "complex-password-123!@#"},
		{"", "password-for-empty-string"},
		{"Special characters: !@#$%^&*()", "another-pwd"},
	}

	for _, tc := range testCases {
		originalData := []byte(tc.originalData)

		encryptedData, err := EncryptData(originalData, tc.password)
		if err != nil {
			t.Fatalf("EncryptData failed: %v", err)
		}

		if bytes.Equal(encryptedData, originalData) && len(originalData) > 0 {
			t.Error("Encrypted data is the same as original data, encryption failed")
		}

		decryptedData, err := DecryptData(encryptedData, tc.password)
		if err != nil {
			t.Fatalf("DecryptData failed: %v", err)
		}

		if !bytes.Equal(decryptedData, originalData) {
			t.Errorf("Decrypted data doesn't match original. Original: %s, Decrypted: %s",
				string(originalData), string(decryptedData))
		}
	}
}

func TestDecryptWithWrongPassword(t *testing.T) {
	originalData := []byte("Secret message")
	correctPassword := "correct-password"
	wrongPassword := "wrong-password"

	encryptedData, err := EncryptData(originalData, correctPassword)
	if err != nil {
		t.Fatalf("EncryptData failed: %v", err)
	}

	_, err = DecryptData(encryptedData, wrongPassword)
	if err == nil {
		t.Error("Decryption with wrong password should fail but succeeded")
	}
} 