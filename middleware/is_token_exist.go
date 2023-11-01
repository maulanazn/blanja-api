package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type AuthenticateTokenHandler func(http.ResponseWriter, *http.Request)

type MakesureToken struct {
	handler AuthenticateTokenHandler
}

func (authenticate *MakesureToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("USR_ID")
	authorization := r.Header.Get("Authorization")
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "Please login correctly")
		return
	}
	if authorization == "" {
		log.Println(err)
		fmt.Fprint(w, "Need Authorization")
		return
	}

	authenticate.handler(w, r)
}

func NewEntranceToken(wrapHandler AuthenticateTokenHandler) *MakesureToken {
	return &MakesureToken{wrapHandler}
}
