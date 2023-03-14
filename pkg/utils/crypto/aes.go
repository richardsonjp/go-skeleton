package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"go-skeleton/pkg/utils/errors"
	"io"
	"strings"

	evp "github.com/walkert/go-evp"
)

var cryptKey = "ac9562b300016a82e9d7e86f0ef8b17a"

func pkcs7Unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func pkcs7Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 8) // Generate an 8 byte salt
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

func DecryptAesWithIv(rawKey string, encryptedData string) (string, error) {
	split := strings.Split(encryptedData, ":")
	if len(split) < 2 {
		return "", errors.NewGenericError(errors.INVALID_ENCRYPTION)
	}
	saltCiphertext, _ := hex.DecodeString(split[1])
	salt := saltCiphertext[8:16]

	// Gets key and IV from raw key.
	key, iv := evp.BytesToKeyAES256CBCMD5([]byte(salt), []byte(rawKey))

	// Creates chipers.
	ciperBlock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Decrypts.
	ciphertext := saltCiphertext[16:]
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(ciperBlock, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	// Unpads.
	unpaddedPlaintext := pkcs7Unpad(plaintext)

	decryptedData := string(unpaddedPlaintext)
	return decryptedData, nil
}

func EncryptAesWithIv(rawKey, text *string) (string, error) {
	// add padding to raw text
	plainText := pkcs7Pad([]byte(*text))

	// Gets key and IV from raw key.
	salt, _ := generateSalt()
	key, iv := evp.BytesToKeyAES256CBCMD5(salt, []byte(*rawKey))

	// Create new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return "", err
	}

	// We need mimic the behavior of CryptoJS format from core
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	copy(cipherText[0:], "Salted__")
	copy(cipherText[8:], salt)

	// Encrypt
	encryptStream := cipher.NewCTR(block, iv)
	encryptStream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return fmt.Sprintf("%x:%x", iv, cipherText), nil
}

func SampleMain() {
	text := "test123123"
	encryptedData, err := EncryptAesWithIv(&cryptKey, &text)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(encryptedData)
	}

	decryptedData, err := DecryptAesWithIv(cryptKey, encryptedData)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(decryptedData)
	}
}
