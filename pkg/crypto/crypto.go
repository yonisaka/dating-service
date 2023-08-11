package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(plainText string) (string, error) {
	textBytes := []byte(plainText)

	hash := sha256.New()

	hash.Write(textBytes)

	hashedTextBytes := hash.Sum(nil)

	hashedText := hex.EncodeToString(hashedTextBytes)

	return hashedText, nil
}

// CompareHash compares a hashed text with a plain text
func CompareHash(hashedText, plainText string) (bool, error) {
	textBytes := []byte(plainText)

	hash := sha256.New()
	hash.Write(textBytes)

	computedHashedTextBytes := hash.Sum(nil)

	computedHashedText := hex.EncodeToString(computedHashedTextBytes)

	// Compare the computed hashed text with the stored hashed text
	return hashedText == computedHashedText, nil
}
