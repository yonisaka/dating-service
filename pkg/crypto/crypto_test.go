package crypto

import "testing"

func TestHashing(t *testing.T) {
	plainText := "password"
	hashedText, err := Hash(plainText)
	if err != nil {
		t.Fatal(err)
	}

	if hashedText == "" {
		t.Fatal("hashedText is empty")
	}

	// Compare the hashed text with the plain text
	ok, err := CompareHash(hashedText, plainText)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("hashedText and plainText do not match")
	}

	t.Log("hashedText and plainText match")
}
