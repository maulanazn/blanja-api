package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"util"
)

type AuthenticateTokenHandler func(http.ResponseWriter, *http.Request)

type MakeSureToken struct {
	handler AuthenticateTokenHandler
}

func (authenticate *MakeSureToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userid, err := r.Cookie("USR_ID")
	authorization := r.Header.Get("Authorization")
  strings.Split(authorization, " ")
  userIdToken := util.DecodeToken(authorization[14:], r)

	if err != nil {
		log.Println(err)
		if _, err := fmt.Fprint(w, "Please login correctly"); err != nil {
			log.Println(err.Error())
		}
		return
	}

	if userid.Value != userIdToken {
		log.Println(err)
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
