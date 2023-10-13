package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var secretKey = GetEnv("SECRET_KEY", "thisis32bitlongpassphraseimusing")

func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(stringToDecrypt string) (string, error) {
	cipherText, err := base64.URLEncoding.DecodeString(stringToDecrypt)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}
