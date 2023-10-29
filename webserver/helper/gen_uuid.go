package helper

import (
	"strings"

	"github.com/google/uuid"
)

func GenUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
