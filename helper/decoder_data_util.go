package helper

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const (
	Passphrase = "abc&1*~#^2^#s0^=)^^7%b34"
)

func DecodeRequest(req *http.Request, data interface{}) error {
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
