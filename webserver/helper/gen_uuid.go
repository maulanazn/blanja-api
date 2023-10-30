package helper

import (
	"strings"

	"github.com/google/uuid"
)

func GenUUID() string {
	uuidWithHyphen := uuid.NewString()
	uuid := strings.Replace(uuidWithHyphen, "-", "", -1)
	return uuid
}
