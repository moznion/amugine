package amugine

import "testing"

func TestEncryptAndDecrypt(t *testing.T) {
	raw := "foobarbuz"
	key := "p@ssW0rd"

	encrypted, err := encrypt([]byte(key), []byte(raw))
	if err != nil {
		t.Fatal("failed to encrypt a payload", err)
	}

	decrypted, err := decrypt([]byte(key), encrypted.Bytes())
	if err != nil {
		t.Fatal("failed to decrypt a payload", err)
	}

	if string(decrypted) != raw {
		t.Error("decrypted payload doesn't equal to original value", err)
	}
}
