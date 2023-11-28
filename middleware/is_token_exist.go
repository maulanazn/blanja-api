package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"util"
)

type AuthenticateTokenHandler func(writer http.ResponseWriter, req *http.Request)

type MakeSureToken struct {
	handler AuthenticateTokenHandler
}

func (authenticate *MakeSureToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")
	strings.Split(authorization, " ")
	userIdToken := util.DecodeToken(authorization[14:], r)

	if userIdToken == "" && authorization == "" {
		if _, err := fmt.Fprint(w, "Failed to Authorize"); err != nil {
			log.Println(err.Error())
		}
		return
	}

	authenticate.handler(w, r)
}

func NewEntranceToken(wrapHandler AuthenticateTokenHandler) *MakeSureToken {
	return &MakeSureToken{wrapHandler}
}
