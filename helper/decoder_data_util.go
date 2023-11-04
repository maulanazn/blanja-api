package helper

import (
	"crypto/aes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func DecodeData(req *http.Request, data interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return errors.New("failed to decode")
	}

	return nil
}

func ConvertStrInt64(data interface{}, base int, bitSize int) (int64, error) {
	format, formatErr := strconv.ParseInt(data.(string), base, bitSize)
	if formatErr != nil {
		return 0, formatErr
	}

	return format, nil
}

func Encrypt(text, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(text))
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(secret))
	block.Encrypt(ciphertext, []byte(secret))
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(text, secret string) (string, error) {
	txt, _ := hex.DecodeString(secret)
	c, err := aes.NewCipher([]byte(text))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(txt))
	c.Decrypt(msgByte, []byte(txt))

	msg := string(msgByte[:])
	return msg, nil
}
