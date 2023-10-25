package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

func DecodeData(req *http.Request, data interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(data); err != nil {
		return errors.New("failed to decode")
	}

	return nil
}
